import Vue from 'vue'
Vue.config.devtools = true
import Router from 'vue-router'
Vue.use(Router)

import JobList from '@/components/JobList'

export default new Router({
	mode: 'hash',
	routes: [
		{
			path: '/',
			redirect: '/jobs'
		},
		{
			path: '/jobs',
			name: 'JobList',
			component: JobList
		},
	]
})
