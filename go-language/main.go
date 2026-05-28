package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprint(w, `<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>TopBrand Studio</title>
  <style>
    :root { --bg:#070b12; --bg-soft:#0f1725; --accent:#f4b263; --accent-strong:#f08a2d; --text:#fff9f4; --muted:#b3b9ca; --card:rgba(255,255,255,0.06); --border:rgba(255,255,255,0.1); }
    * { box-sizing: border-box; }
    body { margin:0; font-family: Inter, system-ui, sans-serif; color:var(--text); background:
      radial-gradient(circle at top, rgba(244,178,99,0.18), transparent 25%),
      linear-gradient(180deg, #04070d, #03050a 48%, #010204);
    }
    .wrapper { max-width: 1100px; margin: 0 auto; padding: 22px; }
    .topbar { display:flex; justify-content:space-between; align-items:center; gap:12px; padding:12px 0 24px; }
    .brand { font-size:1.35rem; font-weight:800; }
    .nav a { color:var(--text); text-decoration:none; margin-left:16px; opacity:.8; }
    .hero { display:grid; grid-template-columns:1.2fr .8fr; gap:22px; align-items:center; padding:18px 0 8px; }
    .headline h1 { margin:0; font-size:clamp(2.8rem,6vw,5rem); line-height:.95; }
    .headline p { color:var(--muted); font-size:1.08rem; line-height:1.5; max-width:640px; margin:16px 0 0; }
    .cta-row { display:flex; gap:10px; margin-top:20px; }
    .btn { border-radius:999px; padding:12px 16px; text-decoration:none; font-weight:800; }
    .btn-primary { background:linear-gradient(180deg, var(--accent), var(--accent-strong)); color:#2a1a0f; }
    .btn-secondary { border:1px solid var(--border); color:var(--text); background:rgba(255,255,255,0.05); }
    .stat-card, .feature-card, .section-card { background:var(--card); border:1px solid var(--border); border-radius:22px; backdrop-filter: blur(6px); box-shadow: 0 16px 40px rgba(0,0,0,.28); }
    .stat-card { padding:18px; display:flex; flex-direction:column; justify-content:space-between; min-height:240px; }
    .stat-badge { display:inline-block; padding:6px 10px; border-radius:999px; background:rgba(244,178,99,0.18); color:var(--accent); font-size:.62rem; font-weight:800; text-transform:uppercase; letter-spacing:.16em; }
    .stat-card h2 { margin:12px 0 0; font-size:2.1rem; }
    .stat-card p { color:var(--muted); margin:8px 0 0; }
    .metrics { display:grid; grid-template-columns:repeat(2,1fr); gap:10px; margin-top:20px; }
    .metric { background:rgba(255,255,255,0.05); border-radius:16px; padding:12px; }
    .metric strong { display:block; font-size:1.25rem; }
    .section-title { margin:24px 0 10px; color:var(--muted); text-transform:uppercase; letter-spacing:.16em; font-size:.7rem; }
    .feature-grid { display:grid; grid-template-columns:repeat(3,1fr); gap:12px; margin-top:10px; }
    .feature-card { padding:18px; }
    .feature-card h3 { margin:0 0 10px; }
    .feature-card p { margin:0; color:var(--muted); line-height:1.5; }
    .section-card { padding:18px; display:flex; justify-content:space-between; align-items:center; gap:16px; margin-top:12px; }
    .section-card p { margin:0; color:var(--muted); line-height:1.5; }
    @media (max-width:760px) { .hero { grid-template-columns:1fr; } .feature-grid { grid-template-columns:1fr; } }
  </style>
</head>
<body>
  <div class="wrapper">
    <div class="topbar">
      <div class="brand">TopBrand</div>
      <nav class="nav">
        <a href="#services">Services</a>
        <a href="#results">Results</a>
        <a href="#contact">Contact</a>
      </nav>
    </div>

    <section class="hero">
      <div class="headline">
        <p class="stat-badge">Global branding agency</p>
        <h1>Build the brand everyone remembers.</h1>
        <p>TopBrand helps ambitious companies launch campaigns, sharpen identity, and grow demand with elegant creative systems and conversion-focused messaging.</p>
        <div class="cta-row">
          <a class="btn btn-primary" href="#contact">Book a strategy call</a>
          <a class="btn btn-secondary" href="#services">See services</a>
        </div>
      </div>

      <div class="stat-card">
        <div>
          <p class="stat-badge">Performance snapshot</p>
          <h2>4.8x</h2>
          <p>Average campaign lift across launch, paid social, and content systems built for scale.</p>
        </div>
        <div class="metrics">
          <div class="metric"><strong>32</strong><span>markets launched</span></div>
          <div class="metric"><strong>98%</strong><span>client retention</span></div>
        </div>
      </div>
    </section>

    <div class="section-title" id="services">Creative services</div>
    <section class="feature-grid">
      <article class="feature-card">
        <h3>Brand Systems</h3>
        <p>Identity direction, messaging frameworks, and launch assets that make bold brands feel premium.</p>
      </article>
      <article class="feature-card">
        <h3>Campaign Design</h3>
        <p>High-velocity creative production for product reveals, events, social, and paid media.</p>
      </article>
      <article class="feature-card">
        <h3>Growth Storytelling</h3>
        <p>Conversion-led copy and content strategy that turns attention into loyal buyers.</p>
      </article>
    </section>

    <div class="section-title" id="results">Trusted by leaders</div>
    <section class="section-card">
      <div>
        <h3 style="margin:0 0 8px;">Headline campaigns with measurable lift</h3>
        <p>We partner with fast-moving startups and established names to shape launches that earn press, clicks, and customer loyalty.</p>
      </div>
      <div class="metric" style="min-width:160px; margin:0;"><strong>$21M</strong><span>tracked revenue influence</span></div>
    </section>

    <div class="section-title" id="contact">Start your next growth chapter</div>
    <section class="section-card">
      <div>
        <h3 style="margin:0 0 8px;">hello@topbrand.studio</h3>
        <p>Creative strategy, campaign direction, and launch design for ambitious companies ready to stand out.</p>
      </div>
      <a class="btn btn-primary" href="mailto:hello@topbrand.studio">Contact studio</a>
    </section>
  </div>
</body>
</html>
`)
    })

    fmt.Println("Go website running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
