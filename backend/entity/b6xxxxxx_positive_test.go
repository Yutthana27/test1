package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	// กรณีข้อมูลถูกต้องครบถ้วน
	// Name: มีค่า
	// CustomerID: ขึ้นต้นด้วย L ตามด้วยเลข 7 ตัว (ถูกต้องตาม Format)
	customer := Customer{
		Name:       "Yutthana",
		Email:      "test@example.com",
		CustomerID: "L1234567",
	}

	// ตรวจสอบ
	ok, err := govalidator.ValidateStruct(customer)

	// คาดหวังว่า ok ต้องเป็น true และ err ต้องเป็น nil
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}