val rgex = """([KQBRN])([0-9]+)""".r

val t1 ="K4"
t1 match {
  case rgex(p,n) => println(s"$p -> $n")
  case _ => println("error")
}
