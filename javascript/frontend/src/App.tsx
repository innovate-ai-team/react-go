import React, { useEffect, useRef, useState } from 'react'
import { startWebGL } from './WebGLRenderer'
import { connectWS } from './ws'

export default function App() {
  const canvasRef = useRef<HTMLCanvasElement | null>(null)
  const [angle, setAngle] = useState(0)

  useEffect(() => {
    if (!canvasRef.current) return
    const renderer = startWebGL(canvasRef.current)

    const ws = connectWS((msg) => {
      if (msg.type === 'tick' && typeof msg.angle === 'number') {
        setAngle(msg.angle)
        renderer.setAngle(msg.angle)
      }
    })

    return () => {
      ws.close()
      renderer.dispose()
    }
  }, [])

  return (
    <div className="app">
      <h1>React + Go Graphics Engine (demo)</h1>
      <p>Angle from server: {angle.toFixed(2)}</p>
      <canvas ref={canvasRef} width={800} height={600} />
    </div>
  )
}
