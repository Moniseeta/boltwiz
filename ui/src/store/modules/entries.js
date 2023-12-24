import {defineStore} from "pinia";
import axios from "axios";

const BASE_URL = process.env.VUE_APP_API_URL || ''

export default defineStore('entries', {
    state() {
        return {
            entries: [],
            currentStack: [],
            currentPage: 1,
            pageSize: 10,
        }
    },
    actions: {
        async getEntries(request) {
            let url = BASE_URL + '/api/v1/list'
            if (request.filter) {
                url += '?key=' + request.filter
            }

            this.currentStack = request.stack || []

            const response = await axios.post(url, request)
            const entries = response.data.results

            this.currentPage = request.page
            this.pageSize = request.pageSize
            return entries
        },
        async addBuckets(request) {
            await axios.post( BASE_URL + '/api/v1/add_buckets', request)
        },
        async addPairs(request) {
            await axios.post(BASE_URL + '/api/v1/add_pairs', request)
        },
        async renameKey(request) {
            await axios.post(BASE_URL + '/api/v1/rename_key', request)
        },
        async updateValue(request) {
            await axios.post(BASE_URL + '/api/v1/update_value', request)
        },
        async deleteEntry(request) {
            await axios.post(BASE_URL + '/api/v1/delete', request)
        },
    },
    getters: {
        entries() {
            return this.entries
        },
    }
})