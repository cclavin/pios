#!/usr/bin/env bash
set -e

F="C\:/Windows/Fonts/CascadiaCode.ttf"
OUT="/c/Users/cclav/Documents/projects/pios/pios-banner.gif"

# ─── Animation timeline ──────────────────────────────────────────────
# 0.00s  dark background
# 0.10s  "> π PIOS" fades in (0.6s ramp)
# 0.75s  orange accent line snaps in
# 1.20s  tagline fades in (0.5s ramp)
# 1.80s  hold title screen
# 4.00s  three white flash pulses (CRT flicker)
# 4.50s  title elements cut off → terminal box appears
# 4.80s  "$ " prompt appears
# 4.80s  "pios init" types out char-by-char (0.14s/char)
# 5.92s  blinking cursor after last char
# 6.30s  "✓ Templates ejected" appears
# 6.60s  "✓ STATUS.md ready" appears
# 6.90s  "→ Fill min-spec.md to begin" appears
# 7.20s  hold, then GIF loops
# ─────────────────────────────────────────────────────────────────────

FILTER="\
[0:v]\
drawtext=fontfile=${F}:fontsize=58:fontcolor=0xf0883e:text='> π  PIOS':x=(w-text_w)/2:y=34:alpha='(t-0.1)/0.6':enable='lt(t\,4.5)',\
drawbox=x=(w-460)/2:y=99:w=460:h=2:color=0xf0883e:t=fill:enable='between(t\,0.75\,4.5)',\
drawtext=fontfile=${F}:fontsize=15:fontcolor=0x8b949e:text='Contracts over vibes.  Build useful things\, faster.':x=(w-text_w)/2:y=118:alpha='(t-1.2)/0.5':enable='lt(t\,4.5)',\
drawbox=x=0:y=0:w=800:h=200:color=0xffffff@0.12:t=fill:enable='between(t\,4.02\,4.10)',\
drawbox=x=0:y=0:w=800:h=200:color=0xffffff@0.12:t=fill:enable='between(t\,4.20\,4.28)',\
drawbox=x=0:y=0:w=800:h=200:color=0xffffff@0.12:t=fill:enable='between(t\,4.38\,4.46)',\
drawbox=x=80:y=50:w=640:h=115:color=0x161b22:t=fill:enable='gte(t\,4.5)',\
drawbox=x=80:y=50:w=640:h=115:color=0xf0883e@0.45:t=2:enable='gte(t\,4.5)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0x3fb950:text='$ ':x=102:y=70:enable='gte(t\,4.8)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='p':x=126:y=70:enable='between(t\,4.80\,4.94)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='pi':x=126:y=70:enable='between(t\,4.94\,5.08)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='pio':x=126:y=70:enable='between(t\,5.08\,5.22)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='pios':x=126:y=70:enable='between(t\,5.22\,5.36)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='pios ':x=126:y=70:enable='between(t\,5.36\,5.50)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='pios i':x=126:y=70:enable='between(t\,5.50\,5.64)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='pios in':x=126:y=70:enable='between(t\,5.64\,5.78)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='pios ini':x=126:y=70:enable='between(t\,5.78\,5.92)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='pios init':x=126:y=70:enable='gte(t\,5.92)',\
drawtext=fontfile=${F}:fontsize=16:fontcolor=0xe6edf3:text='_':x=218:y=71:enable='gt(mod(t-5.92\,0.7)\,0.35)',\
drawtext=fontfile=${F}:fontsize=13:fontcolor=0x3fb950:text='> Templates ejected':x=102:y=98:enable='gte(t\,6.3)',\
drawtext=fontfile=${F}:fontsize=13:fontcolor=0x3fb950:text='> STATUS.md ready':x=102:y=116:enable='gte(t\,6.6)',\
drawtext=fontfile=${F}:fontsize=13:fontcolor=0xf0883e:text='> Fill min-spec.md to begin':x=102:y=134:enable='gte(t\,6.9)'\
[out]"

echo "→ Rendering pios-banner.gif ..."
ffmpeg -y \
  -f lavfi -i "color=c=0x0d1117:size=800x200:rate=20" \
  -filter_complex "$FILTER" \
  -map "[out]" \
  -t 8.8 -loop 0 "$OUT"

echo ""
echo "✓ Done → $OUT"
ls -lh "$OUT"
