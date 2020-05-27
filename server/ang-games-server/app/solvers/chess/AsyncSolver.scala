package solvers.chess

import java.util.UUID

import akka.actor.ActorSystem
import controllers.PostFormInput
import javax.inject.Inject
import play.api.libs.concurrent.CustomExecutionContext

import scala.collection.mutable
import scala.concurrent.Future


class MyExecutionContext @Inject()(system: ActorSystem)
  extends CustomExecutionContext(system, "my.executor")


case class TaskId(taskId: String)

case class TaskDef(taskId: TaskId, result: Future[Results])

/*class TaskActor extends Actor {
  override def receive: Receive = {
    case  td : TaskDef  =>
      ChessAsyncSolver.taskMap.put(td.taskId, td)
      val results = td.either.left.map(_.solve)

  }
}*/

class ChessAsyncSolver @Inject()(implicit myExecutionContext: MyExecutionContext) {
  def createTask(input: PostFormInput): TaskId = {
    val rgex = """([KQBRN])([0-9]+)""".r
    val pmap = for (ent <- input.pieces) yield ent match {
      case rgex(p, n) =>
        (p -> n.toInt)
      case _ =>
        ("?" -> 0)
    }
    val conf = Config(input.dimension, input.dimension, pmap.toMap)
    val solver = SolverV2(conf)
    val td = TaskDef(TaskId(UUID.randomUUID().toString), Future {
      solver.solve
    })
    ChessAsyncSolver.taskMap.put(td.taskId, td)
    td.taskId
  }
}

object ChessAsyncSolver {
  val taskMap = mutable.Map[TaskId, TaskDef]()
}