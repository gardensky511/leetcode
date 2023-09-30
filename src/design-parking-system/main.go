package design_parking_system

// topic 👉 https://leetcode.com/problems/design-parking-system/

type ParkingSystem struct {
	carType map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{carType: map[int]int{1: big, 2: medium, 3: small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.carType[carType] == 0 {
		return false
	}
	this.carType[carType]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
