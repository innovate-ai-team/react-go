type Renderer = {
  setAngle: (a: number) => void
  dispose: () => void
}

export function startWebGL(canvas: HTMLCanvasElement): Renderer {
  const gl = canvas.getContext('webgl')!

  const vert = gl.createShader(gl.VERTEX_SHADER)!
  gl.shaderSource(vert, `
    attribute vec2 a_pos;
    uniform float u_angle;
    void main(){
      float c = cos(u_angle);
      float s = sin(u_angle);
      mat2 rot = mat2(c, -s, s, c);
      gl_Position = vec4(rot * a_pos, 0.0, 1.0);
    }
  `)
  gl.compileShader(vert)

  const frag = gl.createShader(gl.FRAGMENT_SHADER)!
  gl.shaderSource(frag, `
    precision mediump float;
    void main(){
      gl_FragColor = vec4(0.2, 0.7, 1.0, 1.0);
    }
  `)
  gl.compileShader(frag)

  const prog = gl.createProgram()!
  gl.attachShader(prog, vert)
  gl.attachShader(prog, frag)
  gl.linkProgram(prog)
  gl.useProgram(prog)

  const posLoc = gl.getAttribLocation(prog, 'a_pos')
  const angleLoc = gl.getUniformLocation(prog, 'u_angle')

  const buf = gl.createBuffer()!
  gl.bindBuffer(gl.ARRAY_BUFFER, buf)
  const verts = new Float32Array([
    0, 0.6,
    -0.6, -0.6,
    0.6, -0.6
  ])
  gl.bufferData(gl.ARRAY_BUFFER, verts, gl.STATIC_DRAW)

  gl.enableVertexAttribArray(posLoc)
  gl.vertexAttribPointer(posLoc, 2, gl.FLOAT, false, 0, 0)

  let angle = 0
  let raf = 0

  function render(t: number) {
    gl.viewport(0, 0, gl.canvas.width, gl.canvas.height)
    gl.clearColor(0.03, 0.03, 0.05, 1)
    gl.clear(gl.COLOR_BUFFER_BIT)

    gl.uniform1f(angleLoc, angle)
    gl.drawArrays(gl.TRIANGLES, 0, 3)

    raf = requestAnimationFrame(render)
  }

  raf = requestAnimationFrame(render)

  return {
    setAngle(a: number) { angle = a },
    dispose() { cancelAnimationFrame(raf) }
  }
}
