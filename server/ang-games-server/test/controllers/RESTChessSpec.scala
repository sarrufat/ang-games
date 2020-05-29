package controllers

import org.scalatestplus.play._
import org.scalatestplus.play.guice._
import play.api.libs.json.Json
import play.api.mvc.Result
import play.api.test.CSRFTokenHelper._
import play.api.test.Helpers._
import play.api.test._
import v1.post.PostResource

import scala.concurrent.Future

class RESTChessSpec extends PlaySpec with GuiceOneAppPerTest {

  "PostRouter" should {

    "render ChessPostController.checkComplentio" in {
      val request = FakeRequest(GET, " /v1/games/chess/abcd").withHeaders(HOST -> "localhost:9000").withCSRFToken
      val result :Future[Result] = route(app, request).get
        Json.fromJson[String](contentAsJson(result))
      // val posts: Seq[PostResource] = Json.fromJson[Seq[PostResource]](contentAsJson(home)).get
     //  posts.filter(_.id == "1").head mustBe (PostResource("1","/v1/posts/1", "title 1", "blog post 1" ))
    }

  }

}