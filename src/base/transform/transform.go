package transform

import (
	"math"
	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	translation mgl32.Vec3
	rotation mgl32.Vec3
	scale mgl32.Vec3
	zNear float32;
	zFar float32;
	width float32;
	height float32;
	fov float32;
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

func initPerspective (fov, width, height, zNear, zFar float32) mgl32.Mat4{
	aspectRatio:= width/height
	tanHalfFOV := math.Tan(float64(angleToRadian(fov)/2.0))
	zRange := zNear - zFar

	x1:=float32(1.0/(tanHalfFOV * float64(aspectRatio)))
	x2:=float32(1.0/tanHalfFOV )
	x3:=(-zNear-zFar)/zRange
	x4:=2*zFar *zNear/zRange

	return mgl32.Mat4 { 
		x1,0,0,0,0 , x2,0 ,0,0,0,x3, 1,0,0,x4,0}
}

func (t *Transform) GetProjectedTransformation() mgl32.Mat4 {

	perspectiveMatrix:= initPerspective(t.fov, t.width,t.height, t.zNear, t.zFar);
	transformMatrix := t.GetTransformation()

	return perspectiveMatrix.Mul4(transformMatrix)
}

func (t *Transform) SetProjection(fov, width,height,zNear, zFar float32){
	t.fov = fov
	t.width =width
	t.height = height
	t.zNear=zNear
	t.zFar =zFar
}

func (t *Transform) GetTransformation() mgl32.Mat4 {
	translationMatrix := mgl32.Translate3D(t.translation.X(),t.translation.Y(),t.translation.Z())
	rotationMatrix := mgl32.HomogRotate3DX(angleToRadian(t.rotation.X())).Mul4(mgl32.HomogRotate3DY(angleToRadian(t.rotation.Y()))).Mul4(mgl32.HomogRotate3DZ(angleToRadian(t.rotation.Z())))
	scaleMatrix := mgl32.Scale3D(t.scale.X(),t.scale.Y(),t.scale.Z())
	return translationMatrix.Mul4(rotationMatrix.Mul4(scaleMatrix))
}

func (t *Transform) GetTranslation() mgl32.Vec3 {
	return t.translation;
}

func (t *Transform) SetTranslation(vec mgl32.Vec3) {
	t.translation = vec;
}

func (t *Transform) SetTranslationFull(x, y, z float32) {
	t.translation = mgl32.Vec3{x,y,z};
}

func (t Transform) GetRotation()  mgl32.Vec3 {
	return t.rotation;
}

func (t *Transform) SetRotation(vec mgl32.Vec3){
	t.rotation = vec;
}

func (t *Transform) SetRotationFull(x, y, z float32) {
	t.rotation = mgl32.Vec3{x,y,z};
}


func (t *Transform) GetScale()  mgl32.Vec3 {
	return t.scale;
}

func (t *Transform) SetScale(vec mgl32.Vec3) {
	t.scale = vec;
}

func (t *Transform) SetScaleFull(x, y, z float32)  {
	t.scale = mgl32.Vec3{x,y,z};
}