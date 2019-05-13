package main

import (
	"fmt"
	"math/big"
	"strings"
)

func gcd(x *big.Int, y *big.Int) *big.Int {
	if y.Cmp(big.NewInt(0)) == 0 {
		return x
	} else {
		xmody := new(big.Int)
		xmody = xmody.Mod(x, y)
		return gcd(y, xmody)
	}
}
func f(x *big.Int, n *big.Int) *big.Int {
	xsqrd := x.Mul(x, x)
	plus3 := xsqrd.Add(xsqrd, big.NewInt(int64(3)))
	modn := plus3.Mod(plus3, n)
	return modn
}
func pollardrho(a *big.Int) *big.Int {
	var x = big.NewInt(int64(2))
	var y = big.NewInt(int64(2))
	var d = big.NewInt(int64(1))
	for d.Cmp(big.NewInt(int64(1))) == 0 {
		x := f(x, a)
		y := f(f(y, a), a)
		subbed := x.Sub(x, y)
		abs := subbed.Abs(subbed)
		d = gcd(abs, a)
		// fmt.Println(x, y, d)
		if d.Cmp(big.NewInt(int64(1))) == 1 && a.Cmp(d) == 1 {
			return d
		}
		if d.Cmp(a) == 0 {
			return big.NewInt(int64(-1))
		}
	}
	panic("Panic!!!!")
}

func generateMarcene() []*big.Int {
	//Marcenne number = 2^n - 1
	var s []*big.Int
	for j := 2; j < 1000; j++ {
		newBigInt := big.NewInt(int64(2))
		exp := big.NewInt(int64(j))
		newMarcenne := newBigInt.Exp(newBigInt, exp, nil)
		newMarcenne = newMarcenne.Sub(newMarcenne, big.NewInt(int64(1)))
		s = append(s, newMarcenne)
	}
	return s
}

func testDivisor(testFactor *big.Int, n *big.Int) bool {
	z := new(big.Int)
	// fmt.Println("TestFactor: " + testFactor.String())
	// fmt.Println("N: " + n.String())
	z = z.Rem(n, testFactor)
	if z.Cmp(big.NewInt(int64(0))) == 0 {
		return true
	}
	return false
}

func testDivisorSolution(str string, testFactor *big.Int) (*big.Int, *big.Int) {
	n := new(big.Int)
	n, err := n.SetString(str, 10)
	if !err {
		panic(err)
	}
	if testDivisor(testFactor, n) {
		fmt.Println("Factor Found: " + testFactor.String())
		secondFactor := new(big.Int)
		secondFactor = n.Div(n, testFactor)
		fmt.Println("Second Factor: " + secondFactor.String())
		return testFactor, secondFactor

	} else {
		fmt.Println("Factor solution not found")
		return nil, nil
	}
}

func bruteforceSQRT(str string) *big.Int {
	fmt.Println("Reading in n: " + str)
	pq := new(big.Int)
	sqrt := new(big.Int)
	sqrt, err := sqrt.SetString(str, 10)
	if !err {
		panic(err)
	}
	sqrt.Sqrt(sqrt)
	// If it is even... make odd
	if testDivisor(big.NewInt(int64(2)), sqrt) {
		sqrt = sqrt.Add(sqrt, big.NewInt(int64(1)))
	}

	for {
		pq, err = pq.SetString(str, 10)
		if !err {
			panic(err)
		}

		// fmt.Println("Testing Factor: " + sqrt.String())
		if testDivisor(sqrt, pq) {
			// fmt.Println("Testing Solution: " + sqrt.String())
			return sqrt
		}

		sqrt := sqrt.Sub(sqrt, big.NewInt(int64(2)))

		if sqrt.Cmp(big.NewInt(int64(0))) == 0 {
			return big.NewInt(int64(-1))
		}
	}
	panic("OH SHIT!!! I'm Busted")
}

func calculateTotient(p *big.Int, q *big.Int) *big.Int {
	//φ = (p-1)(q-1)
	p = new(big.Int).Sub(p, big.NewInt(int64(1)))
	q = new(big.Int).Sub(q, big.NewInt(int64(1)))
	return new(big.Int).Mul(p, q)
}

func calculateSecretExponent(e *big.Int, totient *big.Int) *big.Int {
	//d = (1/e) mod φ
	return new(big.Int).ModInverse(e, totient)
}

func egcd(a *big.Int, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(int64(0))) != 0 {
		t := b
		bModa := new(big.Int)
		bModa = bModa.Mod(a, b)
		b = bModa
		a = t
	}
	return a
}

func encrypt(n *big.Int, e *big.Int, message []byte) []byte {
	//c=m^e mod n
	return []byte("This isn't done yet")
}

func decrypt(d *big.Int, n *big.Int, c *big.Int) *big.Int {
	//m=c^d mod n .
	plaintext := new(big.Int)
	plaintext.Exp(c, d, n)
	return plaintext
}

func bigInt() {
	//n and e thanks to rsactf tool
	pStr := "7805622068551395034983074294227914827932592556281432557101799867160043121996329164791493852142033952331091204125384233936237118904494182099698709037828123"
	qStr := "7805622068551395034983074294227914827932592556281432557101799867160043121996329164791493852142033952331091204125384233936237118904494182099698709037828129"
	cStr := "41296290787170212566581926747559000694979534392034439796933335542554551981322424774631715454669002723657175134418412556653226439790475349107756702973735895193117931356004359775501074138668004417061809481535231402802835349794859992556874148430578703014721700812262863679987426564893631600671862958451813895661"
	nStr := "60927735877056559130803069919621859729817223816091468870468728150535102345085544195001142179497747300756976118359991531766104121379004146329976732080428122272205922112100073487631152244297343150154109815442681320311122134731991282281969152492933055882377304091844616671159896354284349735375653609635116671867"
	e := big.NewInt(int64(65537))

	//Factor n to p and q
	// possible := bruteforceSQRT(nStr)
	// p, q := testDivisorSolution(nStr, possible)
	p := new(big.Int)
	p, err := p.SetString(pStr, 10)
	if !err {
		panic(err)
	}

	q := new(big.Int)
	q, err = q.SetString(qStr, 10)
	if !err {
		panic(err)
	}

	var totient = calculateTotient(p, q)
	var d = calculateSecretExponent(e, totient)

	n := new(big.Int)
	n, err = n.SetString(nStr, 10)
	if !err {
		panic(err)
	}

	c := new(big.Int)
	c, err = c.SetString(cStr, 10)
	if !err {
		panic(err)
	}
	message := decrypt(d, n, c)

	//Bytes
	fmt.Println("Decrypt: " + message.String())
}

func base4to10(encodedString string) string {
	return ""
}

func generateSeed(seed string, one string, two string, three string, four string) string {
	seedVal := strings.Replace(seed, one, "0", -1)
	seedVal = strings.Replace(seedVal, two, "1", -1)
	seedVal = strings.Replace(seedVal, three, "2", -1)
	seedVal = strings.Replace(seedVal, four, "3", -1)

	n := new(big.Int)
	n, err := n.SetString(seedVal, 4)
	if !err {
		panic(err)
	}

	// val, err := baseconv.Convert(n.String(), baseconv.DigitsDec, baseconv.Digits64)
	fmt.Println(n.String())
	return n.String()
}

//This was a failed attempt to figure out what the n could be for a crypto challenge once upon a time
func permutations() {
	ebola := "CTGAAATGTTCCGCGAGCCGAACCGATTCACCGCCTAGAAACGTATTGTGCTGGTGTGCGGCGGTTAGAGATATTAGGTAGCGCCGTTACTCTAACATTTCGAATCAACCTTTCAGGGGAGTCACTGCCATCGTAAGTAGAGTACTTAGCATCGATGGCCATGCCTACTAATTACAGGCTGAATGACACTAAACCTTAGTTCACTGACCCGTTTTGTCATGTACTCTTGTGGTATGGGTCTTCAAATTGATCTGATTGGGAAGATAGAAAAACGGCTCTATCCTGGGTCGAGCCTCCCATGAAGCAGTCAAGGGGCCGCGAGGACTTCGATACTTGCCCTGCTCGAGCACATTTTAAAGCTTATTCCACATACTAGACTTACCCCCCGGCGTGTCGTACTGGAAGGTTAAACCTCTTGAGTTGATCTGACAACCTAGACGCGTGCCACGTTGTGTGGGATAGGTCACTCTCATTTCCACGAGGGACCAGAACCTTTGGCAATCCAGTTATTCTGCACTCGTGGCCGCCTCTCCTGGCAGGGGACCGGTAAGTTTGCGTATTCGCCGGGGAGTGGAGACGGATCGTCGTACACTGTTTCGAAAATTTTTGAGGATGGAGAGCAGAGCTATTGGATAAACGCTTGTACAGGTTCAATACTATTAGCAACGTGCCACCGGCACAGCTATCTCTGTTTCGCATGAAAGAGCCGTTAATCACGACGTTTAATCGAAACACATACCGATGGTCTACGAATATTATATCCGATACTAAGTCGGCCGCCGCAGTCCAGACGCCATATCGCTTTGAAGACCCCAAGGCGAACATTAACCGGTACGAGCAACTGCGGAGTGCCCTGCAATAGTCCGTCTGTAAAGGGCCCAGGCTAGGGCAAATAGTCCCTAAAACTAGAGATGGTCAACCGCTATGTGGGGCATTCTCCGTGAGACTCAGCCGTATTACAGTGAGCGTATTCCCAAACTCCCCTTCTGTGTATGACCAGTGTCGCTGCAAATGGACCGAGCAG"
	// G,A,T,C
	generateSeed(ebola, "G", "A", "T", "C")
	generateSeed(ebola, "A", "G", "T", "C")
	generateSeed(ebola, "T", "G", "A", "C")
	generateSeed(ebola, "G", "T", "A", "C")
	generateSeed(ebola, "A", "T", "G", "C")
	generateSeed(ebola, "T", "A", "G", "C")
	generateSeed(ebola, "T", "A", "C", "G")
	generateSeed(ebola, "A", "T", "C", "G")
	generateSeed(ebola, "C", "T", "A", "G")
	generateSeed(ebola, "T", "C", "A", "G")
	generateSeed(ebola, "A", "C", "T", "G")
	generateSeed(ebola, "C", "A", "T", "G")
	generateSeed(ebola, "C", "G", "T", "A")
	generateSeed(ebola, "G", "C", "T", "A")
	generateSeed(ebola, "T", "C", "G", "A")
	generateSeed(ebola, "C", "T", "G", "A")
	generateSeed(ebola, "G", "T", "C", "A")
	generateSeed(ebola, "T", "G", "C", "A")
	generateSeed(ebola, "A", "G", "C", "T")
	generateSeed(ebola, "G", "A", "C", "T")
	generateSeed(ebola, "C", "A", "G", "T")
	generateSeed(ebola, "A", "C", "G", "T")
	generateSeed(ebola, "G", "C", "A", "T")
	generateSeed(ebola, "C", "G", "A", "T")
}

func ebola() {
	permutations()
	// bruteforceSQRT(n.String)
}

func BigIntToStr(bigInt *big.Int) string {
	return fmt.Sprintf("%v", bigInt)
}

func main() {
	// bigInt()
	ebola()
	// mdata, err := asn1.Marshal(65537)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Fatal error1: %s", err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Println(mdata)

	// var n int
	// _, err = asn1.Unmarshal(mdata, &n)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Fatal error2: %s", err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Println("After marshal/unmarshal: ", n)
}
