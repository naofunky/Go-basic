package main

import (
	"fmt"
	// "unsafe"
	// "unsafe"
)

func main() {
	// sl := []int{}

	// sl = append(sl, 1)
	// sl = append(sl, 2)
	// sl = append(sl, 1, 2, 3, 4, 5, 6)

	// for i, v := range sl {
	// 	fmt.Println(
	// 		"インデックス番号", i,
	// 		"value値", v,
	// 	)
	// }

	// var nilSlice []int
	// emptySlice := []int{}

	// fmt.Println(nilSlice == nil)
	// fmt.Println(len(emptySlice) == 0)

	// fmt.Println(len(nilSlice))
	// fmt.Println(len(emptySlice))

	// fmt.Println(cap(nilSlice))
	// fmt.Println(cap(emptySlice))

	// fmt.Println(sl)
	// fmt.Println(sl[0])
	// fmt.Println(len(sl))
	// 最後の要素の取得
	// lastElement := sl[len(sl)-1]
	// fmt.Println(lastElement)

	s4 := make([]int, 0, 4)
	fmt.Println(s4, len(s4), cap(s4))
	s5 := append(s4, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s5, len(s5), cap(s5))
	// s76 := append(s4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	// fmt.Println(s76, len(s76), cap(s76))

	// s5 := make([]int, 3, 6)
	// fmt.Println(s5, len(s5), cap(s5))
	// s6 := s5[1:3]
	// fmt.Println(s6, len(s6), cap(s6))
	// s6[1] = 100
	// fmt.Println(s5, len(s5), cap(s5))
	// fmt.Println(s6, len(s6), cap(s6))

	// s7 := []int{1, 2, 3, 4, 5}
	// s8 := s7[2:5]
	// fmt.Println(s7, len(s7), cap(s7))
	// fmt.Println(s8, len(s8), cap(s8))

	// s9 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// s10 := s9[2:6]
	// fmt.Println(s9, len(s9), cap(s9))
	// fmt.Println(s10, len(s10), cap(s10))

	// s11 := []int{11, 12, 13, 14, 15, 16, 17, 18} // 8,8
	// s12 := s11[2:6]
	// fmt.Println("s11:①", s11, len(s11), cap(s11))
	// fmt.Println("s12:②", s12, len(s12), cap(s12))

	// // スライスの要素を追加した時の要素数と容量
	// s12 = append(s12, 7, 8, 999)
	// fmt.Println("s11:③", s11, len(s11), cap(s11))
	// fmt.Println("s12:④", s12, len(s12), cap(s12))

	// s12[0] = 52
	// fmt.Println("s11:⑤", s11, len(s11), cap(s11))
	// fmt.Println("s12:⑥", s12, len(s12), cap(s12))
	// s11 := []int{11, 12, 13, 14, 15, 16, 17, 18} // 8,8
	// s12 := s11[2:6]
	// fmt.Println("s11:①", s11, len(s11), cap(s11), "ポインタ:", unsafe.Pointer(&s11[0]))
	// fmt.Println("s12:②", s12, len(s12), cap(s12), "ポインタ:", unsafe.Pointer(&s12[0]))

	// // スライスの要素を追加した時の要素数と容量
	// s12 = append(s12, 7, 8, 999)
	// fmt.Println("s11:③", s11, len(s11), cap(s11), "ポインタ:", unsafe.Pointer(&s11[0]))
	// fmt.Println("s12:④", s12, len(s12), cap(s12), "ポインタ:", unsafe.Pointer(&s12[0]))

	// s12[0] = 52
	// fmt.Println("s11:⑤", s11, len(s11), cap(s11), "ポインタ:", unsafe.Pointer(&s11[0]))
	// fmt.Println("s12:⑥", s12, len(s12), cap(s12), "ポインタ:", unsafe.Pointer(&s12[0]))

	// 元のスライスまで変わってしまう
	// s12[0] = 555
	// fmt.Println("s11:", s11, len(s11), cap(s11))
	// fmt.Println("s12:", s12, len(s12), cap(s12))

	// s13 := []int{35, 69, 101, 2, 52, 223, 4035, 22, 76, 99} // 10,10
	// s14 := s13[3:7]
	// fmt.Println(s13, len(s13), cap(s13))
	// fmt.Println(s14, len(s14), cap(s14))

	// 普通の要素の追加
	s17 := []int{}
	// fmt.Println("s17:", s17, len(s17), cap(s17))
	s17 = append(s17, 1, 2, 3)
	// fmt.Println("s17:", s17, len(s17), cap(s17))
	s18 := []int{4, 5, 6}
	s17 = append(s17, s18...) //スリードットをつけると可変調で渡せる
	// fmt.Println("s17:", s17, len(s17), cap(s17))

	// makeメソッドを使ってスライスを作成
	s15 := make([]int, 4, 7)
	s16 := s15[1:3]

	// fmt.Println("s15:", s15, len(s15), cap(s15))
	// fmt.Println("s16:", s16, len(s16), cap(s16))

	// 値を代入して確認
	s16[1] = 100

	// s16にさらに値を追加して確認
	s16 = append(s16, 200)

	// fmt.Println("s15:", s15, len(s15), cap(s15))
	// fmt.Println("s16:", s16, len(s16), cap(s16))

	// copy関数を使ってスライスをコピー
	sc16 := make([]int, len((s15[1:3])))
	// fmt.Println("sc16:", sc16, len(sc16), cap(sc16)) // まずはコピー先のスライス作成
	copy(sc16, s15[1:3]) // copy関数でコピー
	sc16[1] = 2000
	// fmt.Println("s15:", s15, len(s15), cap(s15))
	// fmt.Println("sc16:", sc16, len(sc16), cap(sc16))

	//　コピーする範囲を限定するスライスを複製してデータを変更してみる
	// s30 := make([]int, 4, 6)
	// fs := s30[1:3:3]
	// fmt.Println("s30:", s30, len(s30), cap(s30))
	// fmt.Println("fs:", s30, len(fs), cap(fs))

	// fs[0] = 2
	// fs[1] = 3
	// fmt.Println("s30:", s30, len(s30), cap(s30))
	// fmt.Println("fs:", s30, len(fs), cap(fs))

	// fs = append(fs, 8)
	// fmt.Println("s30:", s30, len(s30), cap(s30))
	// fmt.Println("fs:", s30, len(fs), cap(fs))

	// s30[3] = 999
	// fmt.Println("s30:", s30, len(s30), cap(s30))
	// fmt.Println("fs:", s30, len(fs), cap(fs))
}
