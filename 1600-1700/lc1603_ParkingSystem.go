package main

type ParkingSystem struct {
	cars [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		cars: [3]int{big, medium, small},
	}
}

func (b *ParkingSystem) AddCar(carType int) bool {
	if b.cars[carType-1] <= 0 {
		return false
	}
	b.cars[carType-1]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
