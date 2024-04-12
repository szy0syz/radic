package playground

// IsBit1 判断第i位上是否为1，i从1开始
func IsBit1(n uint64, i int) bool {
	if i > 64 {
		panic(i)
	}
	c := uint64(1 << (i - 1))
	if c&n == c {
		return true
	} else {
		return false
	}
}

// SetBit1 把第i位置为1，i从1开始
func SetBit1(n uint64, i int) uint64 {
	if i > 64 {
		panic(i)
	}
	c := uint64(1 << (i - 1))
	return n | c
}

// CountBit1 一个整数的二进制里包含几个1
func CountBit1(n uint64) int {
	sum := 0
	c := uint64(1)
	for i := 0; i < 64; i++ {
		if c&n == c {
			sum += 1
		}
		// c = c<<1
		c <<= 1
	}

	return sum
}

const (
	MALE = 1 << iota
	VIP
	WEEK_ACTIVE
)

type Candidate struct {
	Id     int
	Gender string
	Vip    bool
	Active int //几天内活跃度
	Bits   uint64
}

func (c *Candidate) SetMale() {
	c.Gender = "男"
	c.Bits |= MALE
}

func (c *Candidate) SetVip() {
	c.Vip = true
	c.Bits |= VIP
}

func (c *Candidate) SetActive(day int) {
	c.Active = day
	if day <= 7 {
		c.Bits |= WEEK_ACTIVE
	}
}

// Filter1 判断3个条件是否同时满足
func (c *Candidate) Filter1(male, vip, weekActive bool) bool {
	if male && c.Gender != "男" {
		return false
	}
	if vip && !c.Vip {
		return false
	}
	if weekActive && c.Active > 7 {
		return false
	}
	return true
}

// Filter2 判断N个条件是否同时满足
func (c Candidate) Filter2(on uint64) bool {
	return c.Bits&on == on
}

type BitMap struct {
	Table uint64
}
