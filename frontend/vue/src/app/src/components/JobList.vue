<template>
	<div>
		<h1>sample</h1>
		<div class="alert alert-warning">
			{{jobs}}
		</div>
	</div>
</template>


<script>
import axios from 'axios'
export default {
	components: {
	},
	data: function () {
		return {
			jobs: "empty",
		}
	},
	mounted () {
		this.getJobs()
	},
	methods: {
		getJobs() {
			axios.get('/api/jobs', {
			})
			.then(response => {
				this.jobs = JSON.stringify(response.data.jobs)
			})
		},
	}
}
</script>


<style>
</style>
