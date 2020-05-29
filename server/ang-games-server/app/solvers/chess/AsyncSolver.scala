package solvers.chess

import java.util.UUID

import akka.actor.ActorSystem
import controllers.PostFormInput
import javax.inject.Inject
import play.api.Logger
import play.api.libs.concurrent.CustomExecutionContext

import scala.collection.mutable
import scala.concurrent.Future


class MyExecutionContext @Inject()(system: ActorSystem)
  extends CustomExecutionContext(system, "my.executor")


case class TaskId(taskId: String)

case class TaskDef(taskId: TaskId, result: Future[RestResult])

/*class TaskActor extends Actor {
  override def receive: Receive = {
    case  td : TaskDef  =>
      ChessAsyncSolver.taskMap.put(td.taskId, td)
      val results = td.either.left.map(_.solve)

  }
}*/

class ChessAsyncSolver @Inject()(implicit myExecutionContext: MyExecutionContext) {
  private val logger = Logger(getClass)

  def createTask(input: PostFormInput): TaskId = {
    val pmap = for (piece <- input.pieces; if piece.npieces > 0) yield (piece.letter -> piece.npieces)
    val dim = input.dim.take(1).toInt
    val conf = Config(dim, dim, pmap.toMap)
    val solver = SolverV2(conf)
    val td = TaskDef(TaskId(UUID.randomUUID().toString), Future {
      solver.solve
    })
    /* td.result.onComplete { results =>
      results match {
        case Success(res) =>
          logger.trace(s"Success onComplete $res")
        case Failure(_) =>
          logger.trace("Success onComplete")
      }
    }*/
    ChessAsyncSolver.taskMap.put(td.taskId.taskId, td)
    td.taskId
  }

  def checkTask(taskId: String): CRestResult = ChessAsyncSolver.taskMap.get(taskId) match {

    case Some(task) =>
      if (task.result.isCompleted) {
        task.result.value match {
          case Some(tr) if tr.isSuccess =>
            RestResult.map(tr.get)
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