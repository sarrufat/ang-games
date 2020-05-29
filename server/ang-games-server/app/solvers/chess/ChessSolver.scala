package solvers.chess

case class CPosRes(piece: String, x: Int, y: Int)

case class CResultPositions(positions: List[CPosRes])

case class CRestResult(done: Boolean, ms: Long, combinations: List[CResultPositions])

case class RestResult(ms: Long, results: Results)

case object RestResult {
  val NoResult = CRestResult(false, 0L, List())

  def map(r: RestResult) = CRestResult(true, r.ms, r.results.map(comb => CResultPositions(comb.map(cp => CPosRes(cp._2.toString, cp._1._1, cp._1._2)))))

}

/**
  * The solver class of Chess Challenge
  */
class SolverV2(dimension: Dimension, pieces: Seq[PieceParam]) {


  val seqPieces = (for {
    p <- pieces
    n <- p._1 to 1 by -1
  } yield p._2).toList.mkString

  assert(dimension._1 > 2 && dimension._2 > 2)

  /**
    * Print result as:
    * +-+-+-+-+
    * |*|R|*|*|
    * +-+-+-+-+
    * |N|*|N|*|
    * +-+-+-+-+
    * |*|*|*|R|
    * +-+-+-+-+
    * |N|*|N|*|
    * +-+-+-+-+
    *
    * @param res
    */
  def printresult(res: ResultPositions) = {
    println("")

    def printlnsep() = {
      val headS = for (x <- 0 until dimension._1) yield "-+"
      println("\n+" + headS.mkString)
    }

    for (y <- 0 until dimension._2) {
      printlnsep
      print('|')
      for (x <- 0 until dimension._1) {
        val pos = (x, y)
        res.find { p => p._1 == pos } match {
          case Some((_, k)) => print(k)
          case None => print('*')
        }
        print('|')
      }
    }
    printlnsep
    println("")
  }

  private def threateningPT(pt: Char): ThreateningVector = {
    var tv = Vector[Vector[Pos]]()
    for {
      x <- 0 until dimension._1
      y <- 0 until dimension._2
    } {
      //      val idx = y * dimension._1 + x
      val board = new Board(dimension._1, dimension._2)
      board.newPiece(pt, (x, y)) match {
        case Some(p) => tv = tv :+ Vector[Pos](p.threatening: _*)
        case None =>
      }
    }
    tv
  }

  /*
   * Pre calculated vectors with the threatening positions by each position of the board and piece type
   */
  private lazy val threateningVectors = {
    // For each piece type al MxN htreatening positions
    (for (pp <- pieces) yield (pp._2 -> threateningPT(pp._2))).toMap
  }

  /**
    * This algorithm solver is a recursive algorithm with
    *
    * @return
    */
  def solve: RestResult = {
    var results: Results = List[ResultPositions]()
    val t0 = System.currentTimeMillis();
    def posToIndex(pos: Pos) = pos._1 * dimension._2 + pos._2

    def recResul(keys: String, thr: Vector[Pos], resPos: ResultPositions): Unit = {

      for {
        x <- 0 until dimension._1
        y <- 0 until dimension._2
        if (!resPos.map(_._1).contains((x, y)) && !thr.contains((x, y)))
      } {
        val k = keys(0)
        val idx = posToIndex(x, y)
        val lastPos = resPos.last
        // Index of last position
        val idxlp = posToIndex(lastPos._1)
        // Skip permutation if already calculated tree
        if (lastPos._2 != k || idx > idxlp) {
          val thrK = threateningVectors.get(k).get(idx)
          // Verify bno threatenin
          val currTree = resPos.map(_._1)
          if (!thrK.exists(currTree.contains(_))) {
            if (keys.length() == 1) results = results :+ ((resPos :+ ((x, y), keys(0))).sortBy {
              _._1
            })
            else recResul(keys.drop(1), thr ++ thrK, resPos ++ List(((x, y), k)))
          }
        }
      }
    }

    val k = seqPieces(0)
    val thrK = threateningVectors.get(k).get
    // The roots of trees
    for {
      x <- 0 until dimension._1
      y <- 0 until dimension._2
    } {
      val idx = posToIndex(x, y)
      recResul(seqPieces.drop(1), thrK(idx), List(((x, y), k)))
    }
    val resMap = results.groupBy {
      _.mkString
    }
    val res = for (m <- resMap) yield m._2.head
    RestResult(System.currentTimeMillis() - t0, res.toList)
  }

  def verboseSolve(print: Boolean, timing: Boolean) = {
    def verbosePieces = {
      val names = Map('K' -> "Kings", 'Q' -> "Queens", 'B' -> "Bishops", 'R' -> "Rooks", 'N' -> "Knights")
      pieces.map { p => s"${p._1} ${names.get(p._2).get} " } mkString (" and ")
    }

    println(s"Trying to solve ${dimension._1}X${dimension._2} board with ${verbosePieces} ...")
    val t0 = System.currentTimeMillis();
    val results = solve
    val t1 = System.currentTimeMillis();
    if (timing) println(s"Found ${results.results.length} solutions in " + (t1 - t0).toDouble / 1000.0 + " secs.")
    else println(s"Found ${results.results.length} solutions")
    if (print) results.results.foreach {
      printresult(_)
    }
  }
}

/**
  * The solver factory object
  */
object SolverV2 {
  /**
    * Creates a solver
    *
    * @param dim dimensions of the board
    * @param pieces
    * @return
    */
  def apply(dim: Dimension)(pieces: PieceParam*): SolverV2 = new SolverV2(dim, pieces)

  /**
    * Creates a solver from configuration
    *
    * @param conf
    * @return
    */
  def apply(conf: Config): SolverV2 = {
    val pieces = conf.pieces.toSeq.map(cp => (cp._2, cp._1.charAt(0))).toSeq
    apply((conf.dimM, conf.dimN))(pieces: _*)
  }
}

/**
  * Configuration class
  *
  */
case class Config(dimM: Int = 4, dimN: Int = 4, pieces: Map[String, Int] = Map().empty, printResult: Boolean = false, timing: Boolean = false)