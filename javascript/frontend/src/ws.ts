type Msg = { type: string; [k: string]: any }

export function connectWS(onMsg: (m: Msg) => void) {
  const url = `ws://localhost:8080/ws`
  const ws = new WebSocket(url)

  ws.onopen = () => {
    console.log('ws open')
  }

  ws.onmessage = (ev) => {
    try {
      const data = JSON.parse(ev.data)
      onMsg(data)
    } catch (e) {
      console.warn('bad ws message', e)
    }
  }

  ws.onclose = () => console.log('ws closed')
  ws.onerror = (e) => console.error('ws error', e)

  return ws
}
