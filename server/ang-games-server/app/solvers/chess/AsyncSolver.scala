package solvers.chess

import java.util.UUID
import java.util.concurrent.TimeoutException

import akka.actor.ActorSystem
import controllers.PostFormInput
import javax.inject.Inject
import play.api.Logger
import play.api.libs.concurrent.CustomExecutionContext

import scala.collection.mutable
import scala.concurrent.Future
import scala.concurrent.duration._
import scala.language.postfixOps
import scala.util.Failure


class MyExecutionContext @Inject()(system: ActorSystem)
  extends CustomExecutionContext(system, "my.executor") {
  def getActorSystem = system
}


case class TaskId(taskId: String)

case class TaskDef(taskId: TaskId, result: Future[(RestResult, String)])


class ChessAsyncSolver @Inject()(implicit myExecutionContext: MyExecutionContext) {
  private val logger = Logger(getClass)

  private def verbosePieces(config: Config) = {
    val names = Map("K" -> "Kings", "Q" -> "Queens", "B" -> "Bishops", "R" -> "Rooks", "N" -> "Knights")
    config.pieces.map { p => s"${p._2} ${names.get(p._1).get} " } mkString (", ")
  }

  private def ellapsedFMT(ellapsed: Long) = ellapsed match {
    case x if (x < 5000L) => s"$ellapsed msecs."
    case x => s"${ellapsed / 1000} secs."
  }

  private def formatMessage(config: Config, results: Results, ellapsed: Long, iterations: Long) = s"Found ${results.length} in ${ellapsedFMT(ellapsed)} and $iterations Iterations for ${config.dimM}X${config.dimN} board with ${verbosePieces(config)}"

  def createTask(input: PostFormInput): TaskId = {
    val pmap = for (piece <- input.pieces; if piece.npieces > 0) yield (piece.letter -> piece.npieces)
    val dim = input.dim.take(1).toInt
    val conf = Config(dim, dim, pmap.toMap)
    val solver = SolverV2(conf)
    val error = akka.pattern.after(60 seconds, myExecutionContext.getActorSystem.getScheduler)(Future.failed[(RestResult, String)](new TimeoutException("Computing too long")))
    val sfuture = Future {
      val result = solver.solve
      val msg = formatMessage(conf, result.results, result.ms, result.iterations)
      (result, msg)
    }

    val td = TaskDef(TaskId(UUID.randomUUID().toString), Future.firstCompletedOf[(RestResult, String)](List(sfuture, error)))
    ChessAsyncSolver.taskMap.put(td.taskId.taskId, td)
    td.taskId
  }

  def checkTask(taskId: String): CRestResult = ChessAsyncSolver.taskMap.get(taskId) match {

    case Some(task) =>
      if (task.result.isCompleted) {
        task.result.value match {
          case Some(tr) =>
            tr match {
              case scala.util.Success(res) =>
                RestResult.map(res)
              case Failure(exception) =>
                CRestResult(true, 0, 0, List(), exception.getMessage)
            }
          case _ =>
            RestResult.NoResult
        }
      } else
        RestResult.NoResult
    case None =>
      RestResult.NoResult
  }
}

object ChessAsyncSolver {
  val taskMap = mutable.Map[String, TaskDef]()
}