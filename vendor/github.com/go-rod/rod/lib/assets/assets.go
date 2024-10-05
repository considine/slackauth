// Package assets is generated by "lib/assets/generate"
package assets

// MousePointer for rod.
const MousePointer = `<?xml version="1.0" encoding="UTF-8"?>
<svg width="277px" height="401px" viewBox="0 0 277 401" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
    <!-- Generator: Sketch 52.6 (67491) - http://www.bohemiancoding.com/sketch -->
    <title>mouse-pointer</title>
    <desc>Created with Sketch.</desc>
    <defs>
        <polygon id="path-1" points="0 0 0 299 60 241 103 341 170 313 130 218 217 216"></polygon>
        <filter x="-24.2%" y="-11.0%" width="148.4%" height="130.8%" filterUnits="objectBoundingBox" id="filter-2">
            <feOffset dx="0" dy="15" in="SourceAlpha" result="shadowOffsetOuter1"></feOffset>
            <feGaussianBlur stdDeviation="15" in="shadowOffsetOuter1" result="shadowBlurOuter1"></feGaussianBlur>
            <feColorMatrix values="0 0 0 0 0.138818027   0 0 0 0 0.138818027   0 0 0 0 0.138818027  0 0 0 0.502660779 0" type="matrix" in="shadowBlurOuter1"></feColorMatrix>
        </filter>
    </defs>
    <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
        <g id="mouse-pointer" transform="translate(30.000000, 15.000000)">
            <g id="outside">
                <use fill="black" fill-opacity="1" filter="url(#filter-2)" xlink:href="#path-1"></use>
                <use fill="#FFFFFF" fill-rule="evenodd" xlink:href="#path-1"></use>
            </g>
            <polygon id="inside" fill="#000000" points="18 44 18 255 66 207 110 313 145 299 102 198 171 197"></polygon>
        </g>
    </g>
</svg>`

// Monitor for rod.
const Monitor = `<html>
  <head>
    <title>Rod Monitor - Pages</title>
    <style>
      body {
        margin: 0;
        background: #2d2c2f;
        color: white;
        padding: 20px;
        font-family: sans-serif;
      }
      a {
        color: white;
        padding: 1em;
        margin: 0.5em 0;
        font-size: 1em;
        text-decoration: none;
        display: block;
        border-radius: 0.3em;
        border: 1px solid transparent;
        background: #212225;
      }
      a:visited {
        color: #c3c3c3;
      }
      a:hover {
        background: #25272d;
        border-color: #8d8d96;
      }
    </style>
  </head>
  <body>
    <h3>Choose a Page to Monitor</h3>

    <div id="targets"></div>

    <script>
      async function update() {
        const list = await (await fetch('/api/pages')).json()
        let html = ''
        list.forEach((el) => {
          html += ` + "`" + `<a href='/page/${el.targetId}' title="${el.url}">${el.title}</a>` + "`" + `
        })

        window.targets.innerHTML = html

        setTimeout(update, 1000)
      }

      update()
    </script>
  </body>
</html>
`

// MonitorPage for rod.
const MonitorPage = `<html>
  <head>
    <style>
      body {
        margin: 0;
        background: #2d2c2f;
        color: #ffffff;
      }
      .navbar {
        font-family: sans-serif;
        border-bottom: 1px solid #1413158c;
        display: flex;
        flex-direction: row;
      }
      .error {
        color: #ff3f3f;
        background: #3e1f1f;
        border-bottom: 1px solid #1413158c;
        display: none;
        padding: 10px;
        margin: 0;
      }
      input {
        background: transparent;
        color: white;
        border: none;
        border: 1px solid #4f475a;
        border-radius: 3px;
        padding: 5px;
        margin: 5px;
      }
      .title {
        flex: 2;
      }
      .url {
        flex: 5;
      }
      .rate {
        flex: 1;
      }
    </style>
  </head>
  <body>
    <div class="navbar">
      <input
        type="text"
        class="title"
        title="title of the remote page"
        readonly
      />
      <input type="text" class="url" title="url of the remote page" readonly />
      <input
        type="number"
        class="rate"
        value="0.5"
        min="0"
        step="0.1"
        title="refresh rate (second)"
      />
    </div>
    <pre class="error"></pre>
    <img class="screen" />
  </body>
  <script>
    const id = location.pathname.split('/').slice(-1)[0]
    const elImg = document.querySelector('.screen')
    const elTitle = document.querySelector('.title')
    const elUrl = document.querySelector('.url')
    const elRate = document.querySelector('.rate')
    const elErr = document.querySelector('.error')

    document.title = ` + "`" + `Rod Monitor - ${id}` + "`" + `

    async function update() {
      const res = await fetch(` + "`" + `/api/page/${id}` + "`" + `)
      const info = await res.json()
      elTitle.value = info.title
      elUrl.value = info.url

      await new Promise((resolve, reject) => {
        const now = new Date()
        elImg.src = ` + "`" + `/screenshot/${id}?t=${now.getTime()}` + "`" + `
        elImg.style.maxWidth = innerWidth + 'px'
        elImg.onload = resolve
        elImg.onerror = () => reject(new Error('error loading screenshots'))
      })
    }

    async function mainLoop() {
      try {
        await update()
        elErr.attributeStyleMap.delete('display')
      } catch (err) {
        elErr.style.display = 'block'
        elErr.textContent = err + ''
      }

      setTimeout(mainLoop, parseFloat(elRate.value) * 1000)
    }

    mainLoop()
  </script>
</html>
`
