package controllers

import javax.inject.Inject
import play.api.Logger
import play.api.libs.json.{JsValue, _}
import play.api.mvc.Action
import solvers.chess.{ChessAsyncSolver, TaskId}

import scala.concurrent.ExecutionContext

case class PostFormInput(dimension: Int, pieces: Array[String])

object PostFormInput {
  implicit val formInputRead = Json.reads[PostFormInput]
}

class ChessPostController @Inject()(cc: PostControllerComponents, solver: ChessAsyncSolver)(
  implicit ec: ExecutionContext)
  extends PostBaseController(cc) {
  private val logger = Logger(getClass)


  def solveGame: Action[JsValue] = Action(parse.json) { request =>
      implicit  val writer = Json.writes[TaskId]
     Json.fromJson[PostFormInput](request.body) match {
      case JsSuccess(input, _) =>
        val taksId = solver.createTask(input)
        Ok(Json.toJson[TaskId](taksId) )
      case e @ JsError(_) =>
        BadRequest("JsError")
    }
  }

  /* private def processJsonPost[A]()(
     implicit request: PostRequest[A]): Future[Result] = {
     def failure(badForm: Form[PostFormInput]) = {
       Future.successful(BadRequest(badForm.errorsAsJson))
     }

     def success(input: PostFormInput) = {
       postResourceHandler.create(input).map { post =>
         Created(Json.toJson(post)).withHeaders(LOCATION -> post.link)
       }
     }

     form.bindFromRequest().fold(failure, success)
   } */
}