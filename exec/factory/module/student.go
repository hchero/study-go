package module

type student struct {
	name  string
	score float32
}

func NewStudent(name string, score float32) student {
	var stu student
	stu.name = name
	stu.score = score
	return stu
}
