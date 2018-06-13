package camera


import (
	"math"
	"github.com/go-gl/mathgl/mgl32"
)

var Y_AXIS = mgl32.Vec3{0,1,0};

type Camera struct {
	position mgl32.Vec3
	forward mgl32.Vec3
	up mgl32.Vec3
};

func NewCamera() (Camera) {
	return Camera{ position: mgl32.Vec3{0,0,0}.Normalize() , forward: mgl32.Vec3{0,0,1}.Normalize(), up: mgl32.Vec3{0,1,0}.Normalize()};
}

func (c *Camera) GetPosition() (mgl32.Vec3){
	return c.position
}

func (c *Camera) GetForward() (mgl32.Vec3){
	return c.forward
}

func (c *Camera) GetUp() (mgl32.Vec3){
	return c.up
}

func angleToRadian( angle float32) float32{
	return math.Pi*angle/180.0
}
func  Rotate(angle float32, axis mgl32.Vec3) (res mgl32.Vec3){
	// sinHalfAngle:= math.Sin(float64(angleToRadian(float32(angle/2.0))))
	// cosHalfAngle:= math.Cos(float64(angleToRadian(angle/2.0)))

	// rX := axis.X()
	// rY := axis.Y()
	// rZ := axis.Z()
	// rW := 1
	return 

}

func (c *Camera) RotateX(angle float32)  {
	Haxis := Y_AXIS.Cross(c.forward).Normalize()
	
	//c.forward= c.forward.Rotate(angle, Y_AXIS).Normalize();

	c.up = c.forward.Cross(Haxis).Normalize()
}

func (c *Camera) RotateY(angle float32)  {
	Haxis := Y_AXIS.Cross(c.forward).Normalize()
	//c.forward= c.forward.Rotate(angle, Haxis).Normalize();

	c.up = c.forward.Cross(Haxis).Normalize()
}

func (c *Camera) Move(dir mgl32.Vec3, amt float32){
	c.position = c.position.Add(dir.Mul(amt))
}

func (c *Camera) GetLeft()  (left mgl32.Vec3){
	left = c.up.Cross(c.forward).Normalize()
	return 
}

func (c *Camera) GetRight()  (left mgl32.Vec3){
	left = c.forward.Cross(c.up).Normalize()
	return 
}