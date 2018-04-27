#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec3 vert;
in vec2 vertTexCoord;
out vec3 vertexCoord;
void main() {
    
    gl_Position = projection * camera * model * vec4(vert, 1);
    vertexCoord = vert;
}