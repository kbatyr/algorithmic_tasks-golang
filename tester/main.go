package main

//////////////////////////////  Atoi  ///////////////////////////////////

// func main() {
// 	fmt.Println(student.Atoi("12345"))
// 	fmt.Println(student.Atoi("0000000012345"))
// 	fmt.Println(student.Atoi("012 345"))
// 	fmt.Println(student.Atoi("Hello World!"))
// 	fmt.Println(student.Atoi("+1234"))
// 	fmt.Println(student.Atoi("-1234"))
// 	fmt.Println(student.Atoi("++1234"))
// 	fmt.Println(student.Atoi("--1234"))
// }

// $ go run .
// 12345
// 12345
// 0
// 0
// 1234
// -1234
// 0
// 0
// $

//////////////////////////////  RecursivePower  ///////////////////////////////////

// func main() {
// 	arg1 := 4
// 	arg2 := 3
// 	fmt.Println(student.RecursivePower(arg1, arg2))
// }

// $ go run .
// 64
// $

//////////////////////////////  PrintCombN  ///////////////////////////////////

// func main() {
// 	student.PrintCombN(1)
// 	student.PrintCombN(3)
// 	student.PrintCombN(9)
// }

// $ go run .
// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
// 012, 013, 014, 015, 016, 017, 018, ... 679, 689, 789
// 012345678, 012345679, ..., 123456789
// $

//////////////////////////////  PrintNbrBase  ///////////////////////////////////

// func main() {
// 	student.PrintNbrBase(125, "0123456789") // 125
// 	z01.PrintRune('\n')
// 	student.PrintNbrBase(-125, "01") // -1111101
// 	z01.PrintRune('\n')
// 	student.PrintNbrBase(125, "0123456789ABCDEF") // 7D
// 	z01.PrintRune('\n')
// 	student.PrintNbrBase(-125, "choumi") // -uoi
// 	z01.PrintRune('\n')
// 	student.PrintNbrBase(125, "aa") // NV
// 	z01.PrintRune('\n')
// }

//////////////////////////////  AtoiBase  ///////////////////////////////////

// func main() {
// 	fmt.Println(student.AtoiBase("125", "0123456789")) // 125
// 	fmt.Println(student.AtoiBase("1111101", "01")) // 125
// 	fmt.Println(student.AtoiBase("7D", "0123456789ABCDEF")) // 125
// 	fmt.Println(student.AtoiBase("uoi", "choumi")) // 125
// 	fmt.Println(student.AtoiBase("bbbbbab", "-ab")) // 0
// }

//////////////////////////////  SplitWhiteSpaces  ///////////////////////////////////

// func main() {
// 	fmt.Println(student.SplitWhiteSpaces("Hello how are you?"))
// 	fmt.Println(student.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"))
// 	fmt.Println(student.SplitWhiteSpaces("Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"))
// 	fmt.Println(student.SplitWhiteSpaces("Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."))
// 	fmt.Println(student.SplitWhiteSpaces("Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"))
// 	fmt.Println(student.SplitWhiteSpaces("In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"))
// }

//////////////////////////////  Split  ///////////////////////////////////

// func main() {
// 	s := "HelloHAhowHAareHAyou?"
// 	fmt.Println(student.Split(s, "HA"))
// 	s = "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|" +
// 		"central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished"
// 	fmt.Println(student.Split(s, "|=choumi=|"))
// 	s = "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was" +
// 		"!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator"
// 	fmt.Println(student.Split(s, "!==!"))
// 	s = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator"
// 	fmt.Println(student.Split(s, "AFJ"))
// 	s = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the" +
// 		"<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he" +
// 		"<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer"
// 	fmt.Println(student.Split(s, "<<==123==>>"))
// }

//////////////////////////////  ConvertBase  ///////////////////////////////////

// func main() {
// 	result := student.ConvertBase("101011", "01", "0123456789") // 43
// 	fmt.Println(result)
// 	result = student.ConvertBase("4506C", "0123456789ABCDEF", "choumi") // hccocimc
// 	fmt.Println(result)
// 	result = student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF") // 9B611
// 	fmt.Println(result)
// 	result = student.ConvertBase("256850", "0123456789", "01") // 111110101101010010
// 	fmt.Println(result)
// 	result = student.ConvertBase("uuhoumo", "choumi", "Zone01") // eeone0n
// 	fmt.Println(result)
// 	result = student.ConvertBase("683241", "0123456789", "0123456789") // 683241
// 	fmt.Println(result)
// }

//////////////////////////////  AdvancedSortWordArr  ///////////////////////////////////

// func main() {
////// a-b //////
// // 1 2 3 A B C a b c
// result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// student.AdvancedSortWordArr(result, student.Compare)
// fmt.Println(result)

// // The computing consisted device each earliest fingers five hand of of the undoubtedly
// result = []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of",
// 	"the", "five", "fingers", "of", "each", "hand"}
// student.AdvancedSortWordArr(result, student.Compare)
// fmt.Println(result)

// // "digits" The comesfrom digital fingers or word
// result = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
// student.AdvancedSortWordArr(result, student.Compare)
// fmt.Println(result)

////// b-a //////
// // undoubtedly the of of hand five fingers earliest each device consisted computing The
// result = []string{"The", "computing", "consisted", "device", "each", "earliest",
// 	"fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
// student.AdvancedSortWordArr(result, student.Compare)
// fmt.Println(result)

// // word or fingers digital comesfrom The "digits"
// result = []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
// student.AdvancedSortWordArr(result, student.Compare)
// fmt.Println(result)

// // c b a C B A 3 2 1
// result = []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
// student.AdvancedSortWordArr(result, student.Compare)
// fmt.Println(result)
// }

//////////////////////////////  ActiveBits  ///////////////////////////////////

// func main() {
// 	nbits := student.ActiveBits(7)
// 	fmt.Println(nbits)
// }

// $ go run .
// 3
// $

//////////////////////////////  SortListInsert  ///////////////////////////////////

// func PrintList(l *student.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func listPushBack(l *student.NodeI, data int) *student.NodeI {
// 	n := &student.NodeI{Data: data}

// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }

// func main() {

// 	var link *student.NodeI

// 	link = listPushBack(link, 1)
// 	link = listPushBack(link, 4)
// 	link = listPushBack(link, 9)

// 	PrintList(link)

// 	link = student.SortListInsert(link, -2)
// 	link = student.SortListInsert(link, 2)
// 	PrintList(link)
// }

// $ go run .
// 1 -> 4 -> 9 -> <nil>
// -2 -> 1 -> 2 -> 4 -> 9 -> <nil>
// $

//////////////////////////////  SortedListMerge  ///////////////////////////////////

// func PrintList(l *student.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func listPushBack(l *student.NodeI, data int) *student.NodeI {
// 	n := &student.NodeI{Data: data}

// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }

// func main() {
// 	var link *student.NodeI
// 	var link2 *student.NodeI

// 	link = listPushBack(link, 3)
// 	link = listPushBack(link, 5)
// 	link = listPushBack(link, 7)

// 	link2 = listPushBack(link2, -2)
// 	link2 = listPushBack(link2, 9)

// 	PrintList(student.SortedListMerge(link2, link))
// }

// $ go run .
// -2 -> 3 -> 5 -> 7 -> 9 -> <nil>
// $

//////////////////////////////  ListRemoveIf  ///////////////////////////////////

// func PrintList(l *student.List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}

// 	fmt.Print(nil, "\n")
// }

// func ListPushBack(l *student.List, data interface{}) {
// 	n := &student.NodeL{Data: data}

// 	if l.Head == nil {
// 		l.Head = n
// 	}
// 	if l.Tail != nil {
// 		l.Tail.Next = n
// 	}
// 	l.Tail = n
// }

// func main() {
// 	link := &student.List{}
// 	link2 := &student.List{}

// 	fmt.Println("----normal state----")
// 	ListPushBack(link2, 1)
// 	PrintList(link2)
// 	student.ListRemoveIf(link2, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link2)
// 	fmt.Println()

// 	fmt.Println("----normal state----")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "There")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "How")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)
// 	PrintList(link)

// 	student.ListRemoveIf(link, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link)
// }

// $ go run .
// ----normal state----
// 1 -> <nil>
// ------answer-----
// <nil>

// ----normal state----
// 1 -> Hello -> 1 -> There -> 1 -> 1 -> How -> 1 -> are -> you -> 1 -> <nil>
// ------answer-----
// Hello -> There -> How -> are -> you -> <nil>
// $

//////////////////////////////  BTreeTransplant  ///////////////////////////////////

// func main() {
// 	root := &student.TreeNode{Data: "4"}
// 	student.BTreeInsertData(root, "1")
// 	student.BTreeInsertData(root, "7")
// 	student.BTreeInsertData(root, "5")
// 	node := student.BTreeSearchItem(root, "1")
// 	replacement := &student.TreeNode{Data: "3"}
// 	root = student.BTreeTransplant(root, node, replacement)
// 	student.BTreeApplyInorder(root, fmt.Println)
// }

// $ go run .
// 3
// 4
// 5
// 7
// $

//////////////////////////////  BTreeApplyByLevel  ///////////////////////////////////

// func main() {
// 	root := &student.TreeNode{Data: "4"}
// 	student.BTreeInsertData(root, "1")
// 	student.BTreeInsertData(root, "7")
// 	student.BTreeInsertData(root, "5")
// 	student.BTreeApplyByLevel(root, fmt.Println)
// }

// $ go run .
// 4
// 1
// 7
// 5
// $

//////////////////////////////  BTreeDeleteNode  ///////////////////////////////////

// func main() {
// 	root := &student.TreeNode{Data: "4"}
// 	student.BTreeInsertData(root, "1")
// 	student.BTreeInsertData(root, "7")
// 	student.BTreeInsertData(root, "5")
// 	node := student.BTreeSearchItem(root, "4")
// 	fmt.Println("Before delete:")
// 	student.BTreeApplyInorder(root, fmt.Println)
// 	root = student.BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	student.BTreeApplyInorder(root, fmt.Println)
// }

// $ go run .
// Before delete:
// 1
// 4
// 5
// 7
// After delete:
// 1
// 5
// 7
// $
