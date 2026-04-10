type Renderer = {
  setAngle: (a: number) => void
  dispose: () => void
}

export function startWebGL(canvas: HTMLCanvasElement): Renderer {
  const gl = canvas.getContext('webgl')!

  const vert = gl.createShader(gl.VERTEX_SHADER)!
  gl.shaderSource(vert, `
    attribute vec3 a_pos;
    uniform mat4 u_model;
    uniform mat4 u_proj;
    void main(){
      gl_Position = u_proj * u_model * vec4(a_pos, 1.0);
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
  const modelLoc = gl.getUniformLocation(prog, 'u_model')
  const projLoc = gl.getUniformLocation(prog, 'u_proj')

  const buf = gl.createBuffer()!
  gl.bindBuffer(gl.ARRAY_BUFFER, buf)
  const verts = new Float32Array([
    0, 0.6, 0,
    -0.6, -0.6, 0,
    0.6, -0.6, 0
  ])
  gl.bufferData(gl.ARRAY_BUFFER, verts, gl.STATIC_DRAW)

  gl.enableVertexAttribArray(posLoc)
  gl.vertexAttribPointer(posLoc, 3, gl.FLOAT, false, 0, 0)

  let angle = 0
  let modelMatrix = new Float32Array([
    1,0,0,0,
    0,1,0,0,
    0,0,1,0,
    0,0,0,1
  ])
  let raf = 0

  function render(t: number) {
    gl.viewport(0, 0, gl.canvas.width, gl.canvas.height)
    gl.clearColor(0.03, 0.03, 0.05, 1)
    gl.clear(gl.COLOR_BUFFER_BIT)
    // set projection (simple orthographic)
    const aspect = gl.canvas.width / gl.canvas.height
    const proj = new Float32Array([
      1/aspect,0,0,0,
      0,1,0,0,
      0,0,1,0,
      0,0,0,1,
    ])
    gl.uniformMatrix4fv(projLoc, false, proj)
    gl.uniformMatrix4fv(modelLoc, false, modelMatrix)
    gl.drawArrays(gl.TRIANGLES, 0, 3)

    raf = requestAnimationFrame(render)
  }

  raf = requestAnimationFrame(render)

  return {
    setAngle(a: number) { angle = a },
    setModelMatrix(m: number[]|Float32Array) { modelMatrix = new Float32Array(m) },
    dispose() { cancelAnimationFrame(raf) }
  }
}
