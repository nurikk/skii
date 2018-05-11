package example
import scala.collection.mutable.ArrayBuffer
import scala.io.Source

class Point(val elevation: Int) {
  private var _isPeak: Boolean = false
  private val _neighbors: ArrayBuffer[Point] = ArrayBuffer()
  private val _elevation = elevation

  def isPeak = _isPeak
  def isPeak_= (newValue: Boolean): Unit = {
    _isPeak = newValue
  }

  def neighbors = _neighbors
  override def toString: String =  {
    val star = if (isPeak) "*" else ""
    _elevation.toString + star
  }
}

object Main {
  var fileName :String = ""
  val directions : List[Array[Int]] = List(
    Array(-1, 0),
    Array(0, -1),
    Array(1, 0),
    Array(0, 1)
  )

  val map : ArrayBuffer[Array[Point]] = ArrayBuffer()
  val peaks : ArrayBuffer[Point] = ArrayBuffer()
  var maxRow: Int = 0
  var maxCell: Int = 0
  var winner : ArrayBuffer[Int] = ArrayBuffer()


  def rebuildMap(): Unit = {
    maxRow = map.length - 1
    maxCell = map(0).length - 1

    map.zipWithIndex foreach { case(row, rowIdx) =>

      row.zipWithIndex.foreach{
        case (cell, cellIdx) => {
          var isPeak = true
          directions.foreach((direction: Array[Int]) => {
            val moveRow  = rowIdx + direction(0)
            val moveCell = cellIdx + direction(1)
            if (0 to maxRow contains(moveRow)) {
              if (0 to maxCell contains(moveCell)) {
                val movePoint = map(moveRow)(moveCell)
                if (cell.elevation > movePoint.elevation) {
                  cell.neighbors.append(movePoint)
                } else {
                  isPeak = false
                }
              }
            }
          })
          cell.isPeak_=(isPeak)
          if (isPeak) {
            peaks.append(cell)
          }
        }
      }
    }
  }
  def getDrop(path: ArrayBuffer[Int]): Int = {
    if (path.length == 0) {
      return  0
    }
    path(0) - path(path.length - 1)
  }
  def findNextSteps(point: Point, currentPath: ArrayBuffer[Int]) :Unit = {
    currentPath.append(point.elevation)
    point.neighbors.foreach((direction: Point) => {
      val copyCurrentPath:ArrayBuffer[Int] = currentPath.clone
      if (direction.neighbors.isEmpty) {
        copyCurrentPath.append(direction.elevation)
        if (copyCurrentPath.length > winner.length || (copyCurrentPath.length == winner.length && getDrop(copyCurrentPath) > getDrop(winner))) {
          winner = copyCurrentPath.clone
        }
      } else {
        findNextSteps(direction, copyCurrentPath)
      }
    })


  }
  def solve(): Unit = {
    peaks.foreach((peak: Point) => {
      findNextSteps(peak, ArrayBuffer())
    })
  }
  def readMap(): Unit = {
    val lines = Source.fromFile(fileName).getLines()

    val pointLines = lines.drop(1)

    for (line <- pointLines) {
      val row = line.split(' ').map(_.toInt).map(elevation => new Point(elevation = elevation))
      map.append(row)
    }
  }
  def parseArgs(args: Array[String]): Unit = {
    if (args.length == 1) {
      fileName = args(0)
    } else {
      args.foreach(println)
      println("Wrong arguments count, expected 1")
      System.exit(1)
    }
  }
  def main(args: Array[String]): Unit = {
    parseArgs(args)
    readMap()
    rebuildMap()
    solve()

    println("Length: " + winner.length)
    println("drop: " + getDrop(winner))
    println(winner.mkString(" "))
  }
}
