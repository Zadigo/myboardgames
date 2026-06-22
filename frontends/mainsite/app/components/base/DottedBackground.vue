<template>
  <div ref="containerRef" :class="['fixed inset-0 overflow-hidden bg-neutral-950', props.class]">
    <canvas ref="canvasRef" class="absolute inset-0 h-full w-full" />

    <!-- Vignette overlay -->
    <div class="pointer-events-none absolute inset-0" :style="{ background: 'radial-gradient(ellipse at center, transparent 0%, transparent 40%, rgba(10,10,10,0.6) 100%)' }" />

    <!-- Content layer -->
    <div v-if="$slots.default" class="relative z-10 h-full w-full">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
interface DotPatternProps {
  class?: string
  dotSize?: number
  gap?: number
  baseColor?: string
  glowColor?: string
  proximity?: number
  glowIntensity?: number
  waveSpeed?: number
}

const props = withDefaults(defineProps<DotPatternProps>(), {
  dotSize: 2,
  gap: 24,
  baseColor: '#404040',
  glowColor: '#22d3ee',
  proximity: 120,
  glowIntensity: 1,
  waveSpeed: 0.5,
})

interface Dot {
  x: number
  y: number
  baseOpacity: number
}

const canvasRef = useTemplateRef<HTMLCanvasElement>('canvasRef')
const containerRef = useTemplateRef<HTMLElement>('containerRef')

const dots = ref<Dot[]>([])
const mouse = ref({ x: -1000, y: -1000 })
let animationId: number | undefined
const startTime = Date.now()

function hexToRgb(hex: string): { r: number, g: number, b: number } {
  const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex)
  return result
    ? {
        r: parseInt(result[1], 16),
        g: parseInt(result[2], 16),
        b: parseInt(result[3], 16),
      }
    : { r: 0, g: 0, b: 0 }
}

const baseRgb = computed(() => hexToRgb(props.baseColor))
const glowRgb = computed(() => hexToRgb(props.glowColor))

function buildGrid() {
  const canvas = canvasRef.value
  const container = containerRef.value
  if (!canvas || !container) return

  const rect = container.getBoundingClientRect()
  const dpr = window.devicePixelRatio || 1

  canvas.width = rect.width * dpr
  canvas.height = rect.height * dpr
  canvas.style.width = `${rect.width}px`
  canvas.style.height = `${rect.height}px`

  const ctx = canvas.getContext('2d')
  if (ctx) ctx.scale(dpr, dpr)

  const cellSize = props.dotSize + props.gap
  const cols = Math.ceil(rect.width / cellSize) + 1
  const rows = Math.ceil(rect.height / cellSize) + 1

  const offsetX = (rect.width - (cols - 1) * cellSize) / 2
  const offsetY = (rect.height - (rows - 1) * cellSize) / 2

  const newDots: Dot[] = []
  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      newDots.push({
        x: offsetX + col * cellSize,
        y: offsetY + row * cellSize,
        baseOpacity: 0.3 + Math.random() * 0.2,
      })
    }
  }
  dots.value = newDots
}

function draw() {
  const canvas = canvasRef.value
  if (!canvas) return

  const ctx = canvas.getContext('2d')
  if (!ctx) return

  const dpr = window.devicePixelRatio || 1
  ctx.clearRect(0, 0, canvas.width / dpr, canvas.height / dpr)

  const { x: mx, y: my } = mouse.value
  const proxSq = props.proximity * props.proximity
  const time = (Date.now() - startTime) * 0.001 * props.waveSpeed
  const base = baseRgb.value
  const glow = glowRgb.value

  for (const dot of dots.value) {
    const dx = dot.x - mx
    const dy = dot.y - my
    const distSq = dx * dx + dy * dy

    const wave = Math.sin(dot.x * 0.02 + dot.y * 0.02 + time) * 0.5 + 0.5
    const waveOpacity = dot.baseOpacity + wave * 0.15
    const waveScale = 1 + wave * 0.2

    let opacity = waveOpacity
    let scale = waveScale
    let r = base.r
    let g = base.g
    let b = base.b
    let glowFactor = 0

    if (distSq < proxSq) {
      const dist = Math.sqrt(distSq)
      const t = 1 - dist / props.proximity
      const easedT = t * t * (3 - 2 * t)

      r = Math.round(base.r + (glow.r - base.r) * easedT)
      g = Math.round(base.g + (glow.g - base.g) * easedT)
      b = Math.round(base.b + (glow.b - base.b) * easedT)

      opacity = Math.min(1, waveOpacity + easedT * 0.7)
      scale = waveScale + easedT * 0.8
      glowFactor = easedT * props.glowIntensity
    }

    const radius = (props.dotSize / 2) * scale

    if (glowFactor > 0) {
      const gradient = ctx.createRadialGradient(dot.x, dot.y, 0, dot.x, dot.y, radius * 4)
      gradient.addColorStop(0, `rgba(${glow.r}, ${glow.g}, ${glow.b}, ${glowFactor * 0.4})`)
      gradient.addColorStop(0.5, `rgba(${glow.r}, ${glow.g}, ${glow.b}, ${glowFactor * 0.1})`)
      gradient.addColorStop(1, `rgba(${glow.r}, ${glow.g}, ${glow.b}, 0)`)
      ctx.beginPath()
      ctx.arc(dot.x, dot.y, radius * 4, 0, Math.PI * 2)
      ctx.fillStyle = gradient
      ctx.fill()
    }

    ctx.beginPath()
    ctx.arc(dot.x, dot.y, radius, 0, Math.PI * 2)
    ctx.fillStyle = `rgba(${r}, ${g}, ${b}, ${opacity})`
    ctx.fill()
  }

  animationId = requestAnimationFrame(draw)
}

function handleMouseMove(e: MouseEvent) {
  const canvas = canvasRef.value
  if (!canvas) return
  const rect = canvas.getBoundingClientRect()
  mouse.value = {
    x: e.clientX - rect.left,
    y: e.clientY - rect.top,
  }
}

function handleMouseLeave() {
  mouse.value = { x: -1000, y: -1000 }
}

let resizeObserver: ResizeObserver | undefined

onMounted(() => {
  buildGrid()
  animationId = requestAnimationFrame(draw)

  const container = containerRef.value
  if (container) {
    resizeObserver = new ResizeObserver(buildGrid)
    resizeObserver.observe(container)
    container.addEventListener('mousemove', handleMouseMove)
    container.addEventListener('mouseleave', handleMouseLeave)
  }
})

onUnmounted(() => {
  if (animationId !== undefined) cancelAnimationFrame(animationId)
  resizeObserver?.disconnect()

  const container = containerRef.value
  if (container) {
    container.removeEventListener('mousemove', handleMouseMove)
    container.removeEventListener('mouseleave', handleMouseLeave)
  }
})

// Rebuild grid when relevant props change
watch(
  () => [props.dotSize, props.gap],
  () => buildGrid()
)
</script>
