package views

import "html/template"

var Page = template.Must(template.New("page").Parse(`
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8"/>
			<link rel="stylesheet" href="/static/css/output.css"/>
			<script src="https://unpkg.com/htmx.org@2.0.4"></script>
		</head>
		<body class="bg-gray-200 dark:bg-gray-950">
			<div class="min-h-screen bg-gray-200 dark:bg-gray-950 transition-colors">
				<nav class="w-full px-6 py-4 flex items-center justify-between bg-red-200 dark:bg-gray-950 border-b border-gray-200 dark:border-gray-800">
					<a href="/" class="flex items-center gap-2 text-gray-900 dark:text-white">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							attr:viewBox="0 0 40 40"
							width="42"
							height="42"
							fill="none"
						>
							<rect x="4" y="12" width="14" height="16" rx="2" stroke="currentColor" stroke-width="3"></rect>
							<rect x="14" y="22" width="14" height="16" rx="2" stroke="currentColor" stroke-width="3"></rect>
						</svg>
						<span class="text-xl font-bold tracking-tight">Ta3alaChain</span>
					</a>
				</nav>
				<section class="w-full px-6 py-20 md:py-32 bg-white dark:bg-gray-950 text-center">
					<div class="max-w-3xl mx-auto space-y-6">
						<h1 class="text-4xl md:text-5xl lg:text-6xl font-extrabold leading-tight text-gray-900 dark:text-white">
							{{.Hero}}
						</h1>
						<p class="text-lg md:text-xl text-gray-500 dark:text-gray-400 max-w-2xl mx-auto">
						  {{.HeroSub}}
							
						</p>
					</div>
				</section>
				<div class="w-full flex justify-center bg-white dark:bg-gray-950">
					<div class="w-2/3 h-px bg-gradient-to-r from-transparent via-gray-300 dark:via-gray-700 to-transparent"></div>
				</div>
				<section class="w-full px-6 py-20 md:py-32 bg-gray-200 dark:bg-gray-900 text-center">
					<div class="max-w-3xl mx-auto space-y-6">
						<script src="https://fast.wistia.com/player.js" async>
            </script>
						<script src="https://fast.wistia.com/embed/mrymffstu8.js" async type="module">
            </script>
						<style>wistia-player[media-id='mrymffstu8']:not(:defined) 
            { background: center / contain no-repeat url('https://fast.wistia.com/embed/medias/mrymffstu8/swatch'); display: block; filter: blur(5px); padding-top:56.25%; }</style>
						<wistia-player media-id="mrymffstu8" aspect="1.7777777777777777"></wistia-player>
					</div>
				</section>
				<div class="w-full flex justify-center bg-white dark:bg-gray-950">
					<div class="w-2/3 h-px bg-gradient-to-r from-transparent via-gray-300 dark:via-gray-700 to-transparent"></div>
				</div>
				<!-- Sobre Nosotros -->
				<section class="w-full px-6 py-20 md:py-32 bg-gray-200 dark:bg-gray-950">
					<div class="max-w-4xl mx-auto flex flex-col md:flex-row items-center gap-10">
						<!-- Photo placeholder -->
						<div class="w-48 h-48 md:w-64 md:h-64 rounded-full bg-gray-300 dark:bg-gray-700 flex-shrink-0 overflow-hidden">
							<!-- Replace with <img src="/static/img/jose.jpg" alt="Jose Cruz" class="w-full h-full object-cover"/> -->
						</div>
						<!-- Text content -->
						<div class="text-center md:text-left space-y-4">
							<h2 class="text-3xl md:text-4xl font-extrabold text-gray-900 dark:text-white">
							  {{.ProfileTitle}}	
							</h2>
							<p class="text-sm font-semibold text-emerald-500 uppercase tracking-wide">
								Lead Blockchain Engineer
							</p>
							<p class="text-base md:text-lg text-gray-600 dark:text-gray-400 leading-relaxed">
							{{.ProfileBody0}} <span class="font-bold text-gray-900 dark:text-white"> {{.ProfileBody1}}</span> {{.ProfileBody2}} <span class="font-bold text-gray-900 dark:text-white">{{.ProfileBody3}}</span>{{.ProfileBody4}}<span class="font-bold text-gray-900 dark:text-white">{{.ProfileBody5}}</span>{{.ProfileBody6}}<span class="font-bold text-gray-900 dark:text-white">{{.ProfileBody7}}</span>{{.ProfileBody8}}<span class="font-bold text-gray-900 dark:text-white">{{.ProfileBody9}}</span>.
							</p>
							<a
								href="https://www.linkedin.com/in/mahmoud-baibars/"
								target="_blank"
								rel="noopener noreferrer"
								class="inline-flex items-center gap-2 text-emerald-500 hover:text-emerald-400 font-semibold transition-colors"
							>
								<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" attr:viewBox="0 0 24 24" fill="currentColor">
									<path
										d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85
                    3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433a2.062 2.062 0 0 1-2.063-2.065 2.064 2.064 0 1 1 2.063 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225
                    0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"
									></path>
								</svg>
								LinkedIn
							</a>
						</div>
					</div>
				</section>
				<!-- Divider -->
				<div class="w-full flex justify-center bg-gray-200 dark:bg-gray-950">
					<div class="w-2/3 h-px bg-gradient-to-r from-transparent via-gray-300 dark:via-gray-700 to-transparent"></div>
				</div>
				<div class="w-full flex justify-center bg-white dark:bg-gray-950">
					<div class="w-2/3 h-px bg-gradient-to-r from-transparent via-gray-300 dark:via-gray-700 to-transparent"></div>
				</div>
				<!-- Lo que obtienes -->
				<section class="w-full px-6 py-20 md:py-32 bg-gray-200 dark:bg-gray-900">
					<div class="max-w-6xl mx-auto">
						<div class="text-center space-y-4">
							<p class="text-sm uppercase tracking-[0.3em] font-semibold text-gray-500 dark:text-gray-400">Lo que obtienes</p>
							<h2 class="text-3xl md:text-5xl font-black text-gray-900 dark:text-white leading-tight">
								Lo que obtienes al unirte a <span class="text-emerald-500">Blockchain Accelerator</span>
							</h2>
							<p class="text-base md:text-lg text-gray-600 dark:text-gray-400 max-w-3xl mx-auto">
								Recreamos el mismo bloque de valor que se destaca en la web oficial: una rejilla con seis beneficios tangibles
								que refuerzan credibilidad y urgencia.
							</p>
						</div>
						<div class="mt-16 grid gap-6 sm:gap-8 sm:grid-cols-2 lg:grid-cols-3">
							<div class="group h-full rounded-2xl border border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900/70 p-8 shadow-lg shadow-emerald-500/5 transition duration-300 hover:-translate-y-1 hover:shadow-emerald-500/30">
								<div class="mb-5 inline-flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 dark:bg-emerald-500/20 text-emerald-600 dark:text-emerald-300">
									<svg xmlns="http://www.w3.org/2000/svg" attr:viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" class="h-6 w-6">
										<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"></path>
									</svg>
								</div>
								<h3 class="text-xl font-semibold text-gray-900 dark:text-white">
									Completarás 20 proyectos reales paso a paso
								</h3>
								<p class="mt-3 text-sm text-gray-600 dark:text-gray-400 leading-relaxed">
									El portfolio reproduce los +20 proyectos DeFi que mencionan y funciona como sustituto de la experiencia laboral.
								</p>
							</div>
							<div class="group h-full rounded-2xl border border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900/70 p-8 shadow-lg shadow-emerald-500/5 transition duration-300 hover:-translate-y-1 hover:shadow-emerald-500/30">
								<div class="mb-5 inline-flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 dark:bg-emerald-500/20 text-emerald-600 dark:text-emerald-300">
									<svg xmlns="http://www.w3.org/2000/svg" attr:viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" class="h-6 w-6">
										<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"></path>
									</svg>
								</div>
								<h3 class="text-xl font-semibold text-gray-900 dark:text-white">Soporte directo 24/7 con Jose Cruz</h3>
								<p class="mt-3 text-sm text-gray-600 dark:text-gray-400 leading-relaxed">
									Sin chatbots: replicamos el mensaje de disponibilidad constante con un copy corto y contundente.
								</p>
							</div>
							<div class="group h-full rounded-2xl border border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900/70 p-8 shadow-lg shadow-emerald-500/5 transition duration-300 hover:-translate-y-1 hover:shadow-emerald-500/30">
								<div class="mb-5 inline-flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 dark:bg-emerald-500/20 text-emerald-600 dark:text-emerald-300">
									<svg xmlns="http://www.w3.org/2000/svg" attr:viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" class="h-6 w-6">
										<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"></path>
									</svg>
								</div>
								<h3 class="text-xl font-semibold text-gray-900 dark:text-white">Acceso a bolsa de empleo +$100k</h3>
								<p class="mt-3 text-sm text-gray-600 dark:text-gray-400 leading-relaxed">
									Conectores directos con recruiters y ofertas internacionales para mantener la promesa salarial.
								</p>
							</div>
							<div class="group h-full rounded-2xl border border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900/70 p-8 shadow-lg shadow-emerald-500/5 transition duration-300 hover:-translate-y-1 hover:shadow-emerald-500/30">
								<div class="mb-5 inline-flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 dark:bg-emerald-500/20 text-emerald-600 dark:text-emerald-300">
									<svg xmlns="http://www.w3.org/2000/svg" attr:viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" class="h-6 w-6">
										<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"></path>
									</svg>
								</div>
								<h3 class="text-xl font-semibold text-gray-900 dark:text-white">Preparación integral de entrevistas</h3>
								<p class="mt-3 text-sm text-gray-600 dark:text-gray-400 leading-relaxed">
									Incluye simulacros técnicos, checklist de negociación y banco de +150 preguntas frecuentes.
								</p>
							</div>
							<div class="group h-full rounded-2xl border border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900/70 p-8 shadow-lg shadow-emerald-500/5 transition duration-300 hover:-translate-y-1 hover:shadow-emerald-500/30">
								<div class="mb-5 inline-flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 dark:bg-emerald-500/20 text-emerald-600 dark:text-emerald-300">
									<svg xmlns="http://www.w3.org/2000/svg" attr:viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" class="h-6 w-6">
										<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"></path>
									</svg>
								</div>
								<h3 class="text-xl font-semibold text-gray-900 dark:text-white">Certificado NFT (ERC-721)</h3>
								<p class="mt-3 text-sm text-gray-600 dark:text-gray-400 leading-relaxed">
									NFT verificable que puedes enlazar a tu wallet y portfolio para validar la formación.
								</p>
							</div>
							<div class="group h-full rounded-2xl border border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900/70 p-8 shadow-lg shadow-emerald-500/5 transition duration-300 hover:-translate-y-1 hover:shadow-emerald-500/30">
								<div class="mb-5 inline-flex h-14 w-14 items-center justify-center rounded-full bg-emerald-100 dark:bg-emerald-500/20 text-emerald-600 dark:text-emerald-300">
									<svg xmlns="http://www.w3.org/2000/svg" attr:viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" class="h-6 w-6">
										<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"></path>
									</svg>
								</div>
								<h3 class="text-xl font-semibold text-gray-900 dark:text-white">Impulso de visibilidad en LinkedIn</h3>
								<p class="mt-3 text-sm text-gray-600 dark:text-gray-400 leading-relaxed">
									Mentoría sobre contenido y growth para llegar a >2.5k conexiones relevantes en Web3.
								</p>
							</div>
						</div>
					</div>
				</section>
			</div>
		</body>
	</html>

`))
