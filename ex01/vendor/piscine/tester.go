package piscine

import "fmt"
import "math/rand"
import "os"
import "strconv"

func output(str string, board []string) {
	fmt.Print(str);
	fmt.Println("=========\n");
	for h := 0; h < len(board); h++{ 
		print(board[h])
		fmt.Println("")
	}
		fmt.Println("\nanswer--\n")
		FOVSecuring(board)
		fmt.Println("")
}

func randFit(max int, frequency int) string {
	//Change the arg os the rand.Intn for the currency of the '.'.
	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < max; i++{
		if rand.Intn(frequency) == 1 {
			str += (string)('0' + rand.Intn(7) + 2)
		} else {
			str += (string)('.')
		}
	}
	return str
}

func CheckGenArg(size, freq, qty int) {
	if size < 2 {
		fmt.Println("Generate Error:\n")
		fmt.Println(">> size is too small. size must be bigger than 2")
	}
	if freq > size * size {
		fmt.Println("Generate Error:\n")
		fmt.Println(">> freq is too big. freq must be less than size * size.")
	}
	if qty <= 0 || qty > 10 {
		fmt.Println("Generate Error:\n")
		fmt.Println(">> qty must be [1 : 10].")
	}
}

func GenerateBoard(s, f, q int) {
	CheckGenArg(s, f, q)
	if s >= 2 {
		fmt.Println("//////////////////////////////////////////");
		fmt.Println("random maps");
		fmt.Println("size : ", s, "freq : ", freq, "qty : ", q)
		fmt.Println("//////////////////////////////////////////");
		for at := 0; at < q; at++ {
			board := []string{}
			for i := 0; i < s; i++{
				board = append(board, "")
			}
			for j := 0; j < s; j++{
				board[j] = randFit(s, f)
			}
			output("rand_test", board)
		}
	}
}
// 	fmt.Println("//////////////////////////////////////////");
// 	fmt.Println("valid maps");
// 	fmt.Println("//////////////////////////////////////////\n");

// 	fmt.Println("\n----success_cases-----");

// 	board := []string{
// 			"...2.",
// 			"..6.4",
// 			"5...6",
// 			"7.6..",
// 			".3...",
// 	}
// 	output("pdf_v1", board)

// 	board = []string{
// 			".....",
// 			"8.8..",
// 			".7.7.",
// 			"..8.5",
// 			".....",
// 	}
// 	output("pdf_v2", board)

// 	board = []string{
// 			"37...",
// 			"..8..",
// 			".....",
// 			"..8..",
// 			"...69",
// 	}
// 	output("pdf_v3", board)


// 	board = []string{
// 			"9....",
// 			".....",
// 			".....",
// 			".....",
// 			".....",
// 	}
// 	output("pdf_v3", board)

// 	fmt.Println("\n----fail_cases-----");

// 	board = []string{
// 			".....",
// 			".....",
// 			"..33.",
// 			"..3..",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			".....",
// 			".....",
// 			".....",
// 			".....",
// 			".....",
// 	}
// 	output("v1", board)

// 	board = []string{
// 			".....",
// 			".....",
// 			".....",
// 			".....",
// 			".....",
// 	}

// 	fmt.Println("//////////////////////////////////////////");
// 	fmt.Println("invalid maps---------");
// 	fmt.Println("//////////////////////////////////////////\n");

// 	board = []string{
// 			"...",
// 	}
// 	output("pdf_inv1", board)

// 	board = []string{
// 			"...A.",
// 			"..6.4",
// 			"5....6",
// 			"7.6..",
// 			".3...",
// 	}
// 	output("pdf_inv2", board)

// 	board = []string{
// 		"...",
// 		"...",
// 		"",
// 	}
// 	output("invs", board)

// 	board = []string{
// 		"....",
// 		"....",
// 		"....",
// 	}
// 	output("invs", board)

// 	board = []string{
// 		"....",
// 		"....",
// 		"....................",
// 	}
// 	output("invs", board)

// 	board = []string{
// 		".......................................",
// 	}
// 	output("invs", board)

// 		board = []string{
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 			"....................",
// 		}
// 	output("invs", board)


// 	board = []string{
// 		".",
// 	}
// 	output("invs?", board)

// 	board = []string{
// 			"BBBBB",
// 			"BBBBB",
// 			"BB9BB",
// 			"BBBBB",
// 			"BBBBB",
// 	}
// 	output("B_FULL", board)

// 	board = []string{
// 			".....",
// 			".....",
// 			"B....",
// 			".....",
// 			".....",
// 	}
// 	output("B_contaminated", board)

// 	board = []string{
// 		"99999",
// 		"99999",
// 		"B9999",
// 		"99999",
// 		"99999",
// 	}
// 	output("B_contaminated", board)

// 	board = []string{
// 		".....",
// 		"...B.",
// 		"B....",
// 		".....",
// 		"...B.",
// 	}
// 	output("B_contaminated2", board)
	
// 	// zero and 1
	
// 	return;
	
// 	fmt.Println("//////////////////////////////////////////");
// 	fmt.Println("random ---------");
// 	fmt.Println("//////////////////////////////////////////\n");
// 	board = []string{
// 			"...1.",
// 			"..6.4",
// 			"5..3.",
// 			"7.6..",
// 			".3...",
// 	}
// 	output("random", board)

// 	board = []string{
// 			".....",
// 			".....",
// 			".....",
// 			".....",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"...2.",
// 			"..6.4",
// 			"5....",
// 			"7.6..",
// 			".3...",
// 	}
// 	output("random", board)

// 	board = []string{
// 			"...1.",
// 			"..6.4",
// 			"5....6",
// 			"7.6..",
// 			".3...",
// 	}
// 	output("random", board)

// 	board = []string{
// 			"...3.",
// 			"..6.4",
// 			"5....6",
// 			"7.6..",
// 			".3...",
// 	}
// 	output("random", board)

// 	board = []string{
// 			"...6.",
// 			"..6.4",
// 			"5....6",
// 			"7.6..",
// 			".3...",
// 	}
// 	output("random", board)

// 	board = []string{
// 			"...1.",
// 			"..6.4",
// 			"5....",
// 			"7.6..",
// 			".3...",
// 	}
// 	output("random", board)

// 	board = []string{
// 			".....",
// 			".....",
// 			"..3..",
// 			".....",
// 			".....",
// 	}
// 	output("empty", board)
// 	board = []string{
// 			".....",
// 			".....",
// 			"..3..",
// 			"...5.",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			".....",
// 			".22..",
// 			".....",
// 			".....",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			".....",
// 			".2...",
// 			".....",
// 			".....",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			".....",
// 			".3.3.",
// 			".....",
// 			".4.3.",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			".....",
// 			".33..",
// 			".....",
// 			".....",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			".....",
// 			".88..",
// 			".....",
// 			".....",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			".....",
// 			".4.4.",
// 			".4.4.",
// 			".....",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"5....",
// 			".5...",
// 			"..5..",
// 			"...5.",
// 			"....5",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"6....",
// 			".....",
// 			".6...",
// 			"...6.",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"7....",
// 			".7...",
// 			"..7..",
// 			"...7.",
// 			"....7",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			".3...",
// 			"..3..",
// 			"..3..",
// 			"..3..",
// 			"..3..",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"..3..",
// 			"..3..",
// 			"..3..",
// 			"..3..",
// 			"..3..",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"..5..",
// 			"..5..",
// 			"..5..",
// 			"..5..",
// 			"..5..",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"...3.",
// 			".5.3.",
// 			".....",
// 			".2...",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"...4.",
// 			".5...",
// 			".....",
// 			".37..",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"3...5",
// 			".....",
// 			".....",
// 			".....",
// 			".....",
// 	}
// 	output("empty", board)

// 	board = []string{
// 			"99999",
// 			"99999",
// 			"99999",
// 			"99999",
// 			"99999",
// 	}
// 	output("empty", board)
// }