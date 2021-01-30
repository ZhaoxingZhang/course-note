package bits

/*
 * bitXor - x^y using only ~ and &
 *   Example: bitXor(4, 5) = 1
 *   Legal ops: ~ &
 *   Max ops: 14
 *   Rating: 1
 */
// https://www.geeksforgeeks.org/find-xor-of-two-number-without-using-xor-operator/
func bitXor(x, y int) int {
	//return x^y
	return (x & (^y)) | ((^x )& y)
	//return (^x& ^y) & ^(x& y)
	//return (x | y) & (^x | ^y)
	//return (X || Y) && !(X && Y)
}
/*
 * isTmax - returns 1 if x is the maximum, two's complement number,
 *     and 0 otherwise
 *   Legal ops: ! ~ & ^ | +
 *   Max ops: 10
 *   Rating: 1
 */
func negate(x int) int{
	return ^x+1
}
func tmax() int {
	return int(^uint(0)>>1)
}
/*
 * tmin - return minimum two's complement integer
 *   Legal ops: ! ~ & ^ | + << >>
 *   Max ops: 4
 *   Rating: 1
 */
func tmin() int {
	return negate(tmax()+1) // tmin==~tmax
}
// 2
func isPositive(x int) bool {
	return ((x>>63)&(0x1))!=1
}
/*
首先要构造掩码，使用移位运算符构造出奇数位全1的数mask ，
然后获取输入 x 值的奇数位，其他位清零（mask&x），
然后与 mask 进行异或操作，若相同则最终结果为0，然后返回其值的逻辑非。
*/
/*
 * allOddBits - return 1 if all odd-numbered bits in word set to 1
 *   where bits are numbered from 0 (least significant) to 31 (most significant)
 *   Examples allOddBits(0xFFFFFFFD) = 0, allOddBits(0xAAAAAAAA) = 1
 *   Legal ops: ! ~ & ^ | + << >>
 *   Max ops: 12
 *   Rating: 2
 */
func allOddBits(x int) int{
	mask := 0xAA+(0xAA<<8)
	mask=mask+(mask<<16)
	return bang((mask&x)^mask)
}
/*
 * bang - Compute !x without using !
 *   Examples: bang(3) = 0, bang(0) = 1
 *   Legal ops: ~ & ^ | + << >>
 *   Max ops: 12
 *   Rating: 4
 */
func bang(x int) (y int) {
	//x |= (x>>32)
	x |= (x>>16)
	x |= (x>>8)
	x |= (x>>4)
	x |= (x>>2)
	x |= (x>>1)
	return ^x&0x1;
}
/*
 * isLessOrEqual - if x <= y  then return 1, else return 0
 *   Example: isLessOrEqual(4,5) = 1.
 *   Legal ops: ! ~ & ^ | + << >>
 *   Max ops: 24
 *   Rating: 3
 */
//  int val=!!((x+~y)>>31);
// return (!!x|!y)&((!!x&!y)|(val));
func isLessOrEqual(x, y int) int{
	val :=  bang(bang(((x+ ^y)>>63)))
	x >>= 63
	y >>= 63
	return (x| bang(y))&(x& bang(y)) | (val)
}
func ilog2(x int) (ans int) {
	ans += bang(bang((x>>(32+ans))))<<5
	ans += bang(bang((x>>(16+ans))))<<4
	ans += bang(bang((x>>(8+ans))))<<3
	ans += bang(bang((x>>(4+ans))))<<2
	ans += bang(bang((x>>(2+ans))))<<1 // !0=1, <<1=2
	i := bang(bang((x>>(1+ans))))<<0

	ans += i
	return
}
// 3
/*
 * isAsciiDigit - return 1 if 0x30 <= x <= 0x39 (ASCII codes for characters '0' to '9')
 *   Example: isAsciiDigit(0x35) = 1.
 *            isAsciiDigit(0x3a) = 0.
 *            isAsciiDigit(0x05) = 0.
 *   Legal ops: ! ~ & ^ | + << >>
 *   Max ops: 15
 *   Rating: 3
 */
func isAsciiDigit(x int) int {
	sign := ^(0x1<<63-1)
	upperBound := ^(sign|0x39);
	lowerBound := ^0x30;
	upperBound = sign&(upperBound+x)>>63
	lowerBound = sign&(lowerBound+1+x)>>63
	return bang(upperBound|lowerBound)

}
/*
 * conditional - same as x ? y : z
 *   Example: conditional(2,4,5) = 4
 *   Legal ops: ! ~ & ^ | + << >>
 *   Max ops: 16
 *   Rating: 3
 */
/*
根据 x 的布尔值转换为全0或全1
x==0 时位表示是全0的， x!=0 时位表示是全1的
0的补码是本身，位表示全0；1的补码是-1，位表示全1
*/
func conditional(x, y, z int) int {
	x=bang(x)
	x = ^x+1
	return (x&y)|(^x&z)
}

// 逻辑非就是非 0为1, 非 非0 为0
// 等价于bang
//除了0和最小数（符号位为1，其余为0），
// 其他数都是互为相反数关系（符号位取位或为1）
func logicalNeg(x int) int {
	i := (^x+1) // 右移>>,补符号位
	j := (x|i)
	i = j>>63
	j = i+1
	return j
}
/* howManyBits - return the minimum number of bits required to represent x in
 *             two's complement
 *  Examples: howManyBits(12) = 5
 *            howManyBits(298) = 10
 *            howManyBits(-5) = 4
 *            howManyBits(0)  = 1
 *            howManyBits(-1) = 1
 *            howManyBits(0x80000000) = 32
 */
func howManyBits(x int) {
	//var b16,b8,b4,b2,b1,b0 int
}

// 4 float
func floatScale2(f float32) float32{
	uf := int32(f)
	var c int32
	c = 0x807fffff
	exp := (uf&0x7f800000)>>23
	sign := uf&(int32(0x1)<<31)
	if exp==0 {return float32(uf<<1|sign)}
	if exp==255 {return float32(uf)}
	exp++
	if exp==255 {return float32(0x7f800000|sign)}
	return float32((exp<<23)|(uf&(c)))
}
func floatFloat2Int(f float32) {

}
func floatPower2(x float32) {

}