package app.solvers.chess

import scala.collection.immutable
import scala.collection.mutable.ArrayBuffer
import scala.language.postfixOps

object Board {
  /*
   *
   */
  def genCells(M: Int, N: Int): Seq[(Int, Int)] = for {
    x <- 0 until M
    y <- 0 until N
  } yield (x, y)
}

/**
 * This class represents the board's game
 */

case class Board(M: Int, N: Int) {
  assert(M > 0, "Board dimension must be a positive integer")
  assert(N > 0, "Board dimension must be a positive integer")
  private lazy val allPos = Board.genCells(M, N)
  // Pieces on board by position
  //  var pieces = Map[Pos, Piece]().empty
  val pMatrix: Array[Array[Piece]] = Array.ofDim[Piece](M, N)
  // Tries count
  private var _tryCounter = 0L

  def tryCounter: Long = _tryCounter

  def isInside(pos: Pos): Boolean = pos._1 >= 0 && pos._1 < M && pos._2 >= 0 && pos._2 < N

  def checkFreePos(pos: Pos): Boolean = pMatrix(pos._1)(pos._2) == null && !nonFree.contains(pos)

  private def createNew(pt: Char, pos: Pos): Unit = pt match {
    case 'K' => pMatrix(pos._1)(pos._2) = new King(pos, this)
    case 'B' => pMatrix(pos._1)(pos._2) = new Bishop(pos, this)
    case 'R' => pMatrix(pos._1)(pos._2) = new Rook(pos, this)
    case 'Q' => pMatrix(pos._1)(pos._2) = new Queen(pos, this)
    case 'N' => pMatrix(pos._1)(pos._2) = new Knight(pos, this)
    case _ => throw new Exception(s"Unknown piece type '$pt'")
  }

  /**
   * Creates a new piece on position if possible
   * pt represents the type (K, Q, B R, N)
   */
  def newPiece(pt: Char, pos: Pos): Option[Piece] = {

    if (checkFreePos(pos)) {
      createNew(pt, pos)
      Some(pMatrix(pos._1)(pos._2))
    } else {
      None
    }

  }

  def removePiece(piece: Piece): Unit = pMatrix(piece.pos._1)(piece.pos._2) = null

  def nonFree: immutable.Seq[(Int, Int)] = {
    val nfs = for {
      x <- pMatrix.indices
      y <- pMatrix(x).indices
      if pMatrix(x)(y) != null
    } yield {
      (x, y) +: pMatrix(x)(y).threatening
    }
    nfs.flatten
  }

  // (pieces map { p => (p._1 +: p._2.threatening) }).flatten.toSeq
  def getPossibleCells: List[(Int, Int)] = allPos.filterNot(p => nonFree contains p).toList

  /**
   * Tries to Create a new piece on position if possible and no threatening the other pieces on the board
   */
  def tryNewPiece(pt: Char, pos: Pos): Option[Piece] = {
    _tryCounter += 1
    val np = newPiece(pt, pos)
    np match {
      case Some(curPiece) => if (curPiece.threatening exists { p => pMatrix(p._1)(p._2) != null }) {
        removePiece(curPiece)
        None
      } else {
        np
      }
      case None => None
    }
  }

  def isSolved(ntarg: Int): Boolean = pMatrix.flatten.count(p => p != null) == ntarg

  def toResult: Seq[((Int, Int), Char)] = {
    // Graal mandatory refactoring in order to compile
    val xind = pMatrix.indices.iterator
    var buffer = ArrayBuffer[((Int, Int), Char)]()
    while (xind.hasNext) {
      val x = xind.next()
      val yter = pMatrix(x).indices.iterator
      while (yter.hasNext) {
        val y = yter.next()
        if (pMatrix(x)(y) != null) {
          buffer = buffer :+ ((x, y), pMatrix(x)(y).toChar)
        }
      }
    }
    /*   val nfs = for {
         x <- pMatrix.indices
         y <- pMatrix(x).indices
         if (pMatrix(x)(y) != null)
       } yield {
         ((x, y), pMatrix(x)(y).toChar)
       }
       nfs.toList.sortBy(p => p._1)
       */
    buffer.toList.sortBy(p => p._1)
  }

  override def toString: String = {
    val result = toResult
    var retStr = ""
    result
      .foreach(res => retStr += s"$res ")
    retStr
  }
}

/*
 * A generic piece
 */
sealed trait Piece {
  val pos: Pos
  val board: Board

  /*
   * method for getting the threatening positions on the board
   */
  def threatening: Positions

  /**
   * Recursively compute position increments by vector 'incr'
   */
  final protected def vincr(pos: Pos, incr: Pos): Positions = {
    val newPos = (pos._1 + incr._1, pos._2 + incr._2)
    if (board isInside newPos) newPos +: vincr(newPos, incr)
    else List()
  }

  def toChar: Char
}

/**
 * Base of all pieces
 */
abstract class PieceBase(p: Pos, b: Board) extends Piece {
  val pos: (Int, Int) = p
  val board = b
}

class King(p: Pos, b: Board) extends PieceBase(p, b) {
  private lazy val _threatening = {
    val ret = for {
      x <- -1 to 1
      y <- -1 to 1
    } yield (pos._1 + x, pos._2 + y)
    ret filter { p => p != pos && board.isInside(p) } toList
  }

  def threatening: Positions = _threatening

  def toChar: Char = 'K'
}

/**
 * Trait with Bishop directions
 */
sealed trait BishopMov extends Piece {
  val northEst: Direction = () => vincr(pos, (1, -1))
  val southWest: Direction = () => vincr(pos, (-1, 1))
  val southEst: Direction = () => vincr(pos, (1, 1))
  val northWest: Direction = () => vincr(pos, (-1, -1))
}

class Bishop(p: Pos, b: Board) extends PieceBase(p, b) with BishopMov {
  private lazy val _threatening = northEst() ++ southEst() ++ southWest() ++ northWest()

  def threatening: Positions = _threatening

  def toChar: Char = 'B'
}

/**
 * Trait with Rook directions
 */
sealed trait RookMov extends Piece {
  val north: Direction = () => vincr(pos, (0, -1))
  val south: Direction = () => vincr(pos, (0, 1))
  val est: Direction = () => vincr(pos, (1, 0))
  val west: Direction = () => vincr(pos, (-1, 0))
}

class Rook(p: Pos, b: Board) extends PieceBase(p, b) with RookMov {
  private lazy val _threatening = north() ++ est() ++ south() ++ west()

  def threatening: Positions = _threatening

  def toChar: Char = 'R'
}

class Queen(p: Pos, b: Board) extends PieceBase(p, b) with BishopMov with RookMov {
  private lazy val _threatening = north() ++ northEst() ++ est() ++ southEst() ++ south() ++ southWest() ++ west() ++ northWest()

  def threatening: Positions = _threatening

  def toChar: Char = 'Q'
}

object Knight {
  private lazy val movVectors = List((1, -2), (2, -1), (2, 1), (1, 2), (-1, 2), (-2, 1), (-2, -1), (-1, -2))
}

class Knight(p: Pos, b: Board) extends PieceBase(p, b) {
  private lazy val _threatening = {
    val ret = for {vec <- Knight.movVectors} yield {
      (pos._1 + vec._1, pos._2 + vec._2)
    }
    ret filter { p => board isInside (p) } toList
  }

  def threatening: Positions = _threatening

  def toChar: Char = 'N'
}