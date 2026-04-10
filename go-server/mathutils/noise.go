package mathutils

import "math"

func fract(x float64) float64 { return x - math.Floor(x) }
func smoothstep(t float64) float64 { return t * t * (3 - 2*t) }

func hash3(x, y, z int64) float64 {
  // simple integer hash -> float in [0,1)
  n := x*73856093 ^ y*19349663 ^ z*83492791
  v := math.Sin(float64(n))*43758.5453123
  return fract(v)
}

func lerp(a, b, t float64) float64 { return a + (b-a)*t }

// Noise3 returns a smooth value noise in range [-1,1]
func Noise3(x, y, z float64) float64 {
  xi := int64(math.Floor(x))
  yi := int64(math.Floor(y))
  zi := int64(math.Floor(z))

  xf := fract(x)
  yf := fract(y)
  zf := fract(z)

  // sample corners
  var c [2][2][2]float64
  for i := int64(0); i <= 1; i++ {
    for j := int64(0); j <= 1; j++ {
      for k := int64(0); k <= 1; k++ {
        c[i][j][k] = hash3(xi+i, yi+j, zi+k)*2 - 1
      }
    }
  }

  sx := smoothstep(xf)
  sy := smoothstep(yf)
  sz := smoothstep(zf)

  // trilinear interp
  var i0, i1 float64
  i0 = lerp(lerp(lerp(c[0][0][0], c[1][0][0], sx), lerp(c[0][1][0], c[1][1][0], sx), sy), lerp(lerp(c[0][0][1], c[1][0][1], sx), lerp(c[0][1][1], c[1][1][1], sx), sy), sz)
  i1 = i0

  return i1
}

// FBM provides simple fractal Brownian motion based on Noise3
func FBM3(x, y, z float64, octaves int, lacunarity, gain float64) float64 {
  f := 0.0
  amp := 1.0
  freq := 1.0
  for i := 0; i < octaves; i++ {
    f += Noise3(x*freq, y*freq, z*freq) * amp
    freq *= lacunarity
    amp *= gain
  }
  return f
}
