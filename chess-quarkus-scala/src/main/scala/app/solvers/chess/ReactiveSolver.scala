package app.solvers.chess

import java.util.UUID
import java.util.concurrent.ConcurrentHashMap

import io.vertx.core._
import io.vertx.core.eventbus.Message
import javax.json.bind.JsonbBuilder
import org.jboss.logging.Logger
import org.sarrufat.resource.chess.PostFormInput

import scala.util.{Failure, Success, Try}


object ReactiveSolver {
  val logger: Logger = Logger.getLogger(ReactiveSolver.getClass)

  private val vertx = Vertx.vertx()
  private val executor = vertx.createSharedWorkerExecutor("solver")
  private val taskMap = new ConcurrentHashMap[String, RestResult]()

  private val consumer = vertx.eventBus().consumer("solve", (message: Message[String]) => {
    val jsonb = JsonbBuilder.create()
    val input: PostFormInput = jsonb.fromJson(message.body(), classOf[PostFormInput])
    val pmap = for (piece <- input.pieces; if piece.npieces > 0) yield (piece.letter -> piece.npieces)
    val dim = input.dim.take(1).toInt
    val conf = Config(dim, dim, pmap.toMap)
    val solver = SolverV2(conf)

    executor.executeBlocking((promise: Promise[String]) => {
      Try[RestResult] {
        solver.solve
      } match {
        case Success(result) =>
          logger.info(s"result = ${result.ms}ms.")
          promise.complete("s\"result = ${result.ms}ms.\"")
          taskMap.put(input.taskId, result)
        case Failure(exception) =>
          val ms = exception match {
            case e : SolverV2.MaxResultException => e.ms
            case _ => 0l
          }
          val result = RestResult(ms, 0l)
          logger.info(s"result = ${exception}.")
          promise.complete("s\"result = ${result.ms}ms.\"")
          taskMap.put(input.taskId, result)
      }


    }, (result: AsyncResult[String]) => {

    })

  })

  def send(message: PostFormInput): String = {
    val jsonb = JsonbBuilder.create()
    message.taskId = UUID.randomUUID().toString
    val sm = jsonb.toJson(message)
    vertx.eventBus().send("solve", sm)
    val result = message.taskId
    logger.debug(s"result = $result")
    result
  }

  def checkId(taskId: String): Option[RestResult] = {
      if (taskMap.containsKey(taskId)) {
        Option(taskMap.get(taskId))
      } else {
        None
      }
  }

}


