package transform

import (
	"math"
	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	translation mgl32.Vec3
	rotation mgl32.Vec3
	scale mgl32.Vec3
}


func NewTransform() Transform{
	var t Transform= Transform{}
	t.translation = mgl32.Vec3{0,0,0}
	t.rotation = mgl32.Vec3{0,0,0}
	t.scale = mgl32.Vec3{1,1,1}
	return t;
}

func angleToRadian( angle float32) float32{
	return math.Pi*angle/180.0
}

func (t Transform) GetTransformation() mgl32.Mat4 {
	translationMatrix := mgl32.Translate3D(t.translation.X(),t.translation.Y(),t.translation.Z())
	rotationMatrix := mgl32.HomogRotate3DX(angleToRadian(t.rotation.X())).Mul4(mgl32.HomogRotate3DY(angleToRadian(t.rotation.Y()))).Mul4(mgl32.HomogRotate3DZ(angleToRadian(t.rotation.Z())))
	scaleMatrix := mgl32.Scale3D(t.scale.X(),t.scale.Y(),t.scale.Z())
	return translationMatrix.Mul4(rotationMatrix.Mul4(scaleMatrix))
}

func (t Transform) GetTranslation() mgl32.Vec3 {
	return t.translation;
}

func (t Transform) SetTranslation(vec mgl32.Vec3) Transform {
	t.translation = vec;
	return t
}

func (t Transform) SetTranslationFull(x, y, z float32)  Transform{
	t.translation = mgl32.Vec3{x,y,z};
	return t
}

func (t Transform) GetRotation()  mgl32.Vec3 {
	return t.rotation;
}

func (t Transform) SetRotation(vec mgl32.Vec3) Transform{
	t.rotation = vec;
	return t
}

func (t Transform) SetRotationFull(x, y, z float32)  Transform{
	t.rotation = mgl32.Vec3{x,y,z};
	return t
}


func (t Transform) GetScale()  mgl32.Vec3 {
	return t.scale;
}

func (t Transform) SetScale(vec mgl32.Vec3) Transform{
	t.scale = vec;
	return t
}

func (t Transform) SetScaleFull(x, y, z float32)  Transform{
	t.scale = mgl32.Vec3{x,y,z};
	return t
}