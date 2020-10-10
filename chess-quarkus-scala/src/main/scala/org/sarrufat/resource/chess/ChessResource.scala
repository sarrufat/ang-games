package org.sarrufat.resource.chess

import app.solvers.chess.{ReactiveSolver, RestResult}
import javax.inject.Inject
import javax.ws.rs._
import javax.ws.rs.core.MediaType
import org.jboss.logging.Logger
import org.sarrufat.resource.chess.ChessResource.logger

import scala.beans.BeanProperty


class PieceInput {

  @BeanProperty
  var label: String = ""
  @BeanProperty
  var letter: String = ""
  @BeanProperty
  var npieces: Int = 0


  override def toString = s"PieceInput(label=$label, letter=$letter, npieces=$npieces)"
}

class PostFormInput {

  @BeanProperty
  var dim: String = ""
  @BeanProperty
  var pieces: Array[PieceInput] = Array[PieceInput]()
  @BeanProperty
  var taskId: String = ""

  override def toString = s"PostFormInput(dim=$dim, pieces=[${pieces.mkString(", ")}])"
}


class TaskId {
  @BeanProperty
  var taskId: String = ""

  override def toString = s"TaskId(taskId=$taskId)"
}

class CRPosRes {
  @BeanProperty
  var piece: String = ""
  @BeanProperty
  var x: Int = 0
  @BeanProperty
  var y: Int = 0
}

class CRResultPositions {
  @BeanProperty
  var positions: Array[CRPosRes] = Array[CRPosRes]()
}

class JResult {
  @BeanProperty
  var done: Boolean = false
  @BeanProperty
  var ms: Long = 0
  @BeanProperty
  var iterations: Long = 0
  @BeanProperty
  var combinations: Array[CRResultPositions] = Array[CRResultPositions]()
  @BeanProperty
  var msg: String = ""
}

object JResult {
  def apply(res: RestResult): JResult = {
    val jres = new JResult()
    jres.done = true
    jres.iterations = res.iterations
    jres.ms = res.ms
    jres.combinations = res.results.map { r =>
      val comb = new CRResultPositions
      comb.positions = r.map { cp =>
        val cposres = new CRPosRes
        cposres.piece = cp._2.toString
        cposres.x = cp._1._1
        cposres.y = cp._1._2
        cposres

      }.toArray
      comb
    }.toArray
    jres
  }
}

object ChessResource {
  val logger: Logger = Logger.getLogger(ChessResource.getClass)
}

@Path("/v1/games/chess")
@Produces(Array(MediaType.APPLICATION_JSON))
@Consumes(Array(MediaType.APPLICATION_JSON))
class ChessResource @Inject() ( rsolver: ReactiveSolver) {
   def this() {
     this(null)
   }

  @POST
  def solve(input: PostFormInput): TaskId = {
    logger.debug(s"solve: $input")

    val task = new TaskId
    task.taskId = rsolver.send(input)
    task
  }

  @GET
  @Path("/{taskId}")
  def check(@PathParam("taskId") taskId: String): JResult = {
    rsolver.checkId(taskId) match {
      case Some(r) =>
        JResult(r)
      case None =>
        val res = new JResult
        res
    }
  }
}
