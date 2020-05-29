package controllers

import javax.inject.Inject
import play.api.Logger
import play.api.libs.json.{JsValue, _}
import play.api.mvc.{Action, AnyContent}
import solvers.chess._

import scala.concurrent.ExecutionContext

case class PieceInput(label: String, letter: String, npieces: Int)

case class PostFormInput(dim: String, pieces: Array[PieceInput])

object PostFormInput {
  implicit val formInputRead0 = Json.reads[PieceInput]

  implicit val formInputRead = Json.reads[PostFormInput]
}

class ChessPostController @Inject()(cc: PostControllerComponents, solver: ChessAsyncSolver)(
  implicit ec: ExecutionContext)
  extends PostBaseController(cc) {
  private val logger = Logger(getClass)


  def solveGame: Action[JsValue] = Action(parse.json) { request =>
    implicit val writer = Json.writes[TaskId]
    logger.trace(request.toString())
    logger.trace(request.body.toString())
    Json.fromJson[PostFormInput](request.body) match {
      case JsSuccess(input, _) =>
        val taksId = solver.createTask(input)
        logger.trace(s"JsSuccess $input")
        Ok(Json.toJson[TaskId](taksId))
      case e@JsError(_) =>
        logger.trace("JsError")
        BadRequest("JsError")
    }

  }

  def checkCompletion(id: String): Action[AnyContent] = Action {
    implicit val w0 = Json.writes[CPosRes]
    implicit val w1 = Json.writes[CResultPositions]
    implicit val w2 = Json.writes[CRestResult]
    val result = solver.checkTask(id)
    Ok(Json.toJson[CRestResult](result))
  }
}