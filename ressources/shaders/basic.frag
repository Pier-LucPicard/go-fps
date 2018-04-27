#version 330
uniform sampler2D tex;
in vec3 vertexCoord;
out vec4 outputColor;
void main() {
    outputColor = vec4(vertexCoord.x,vertexCoord.y,vertexCoord.z,1);
}