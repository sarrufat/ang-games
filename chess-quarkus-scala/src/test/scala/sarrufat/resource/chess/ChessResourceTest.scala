package sarrufat.resource.chess

import io.quarkus.test.junit.QuarkusTest
import io.restassured.RestAssured.`given`
import  org.hamcrest.CoreMatchers.`is`
import org.junit.jupiter.api.Test

@QuarkusTest
class ChessResourceTest {
  @Test
  def testSolve() = {
    given()
      .`when`().get("/v1/games/chess")
      .then()
      .statusCode(200)
      .body(`is`("hola"))
  }
}
