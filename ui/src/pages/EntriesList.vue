<template>
  <div style="min-height: calc(100vh - 105px)">
    <div class="row">
      <div class="bg-grey-1 full-width">
        <div class="q-py-xs text-center">
          <div class="flex-inline items-center">
            <div class="row">
              <div class="col-10 offset-1">
                <q-breadcrumbs
                    class="text-grey-5 q-my-sm text-orange"
                    active-color="primary"
                    separator="/"
                >
                  <q-breadcrumbs-el label="Home" icon="home" to="/"/>
                  <template
                      :key="item"
                      v-for="(item, index) in stack">
                    <q-breadcrumbs-el
                        v-if="index < stack.length - 1"
                        :label="item"
                        :to="{ name: 'home', params: { stack: stack.slice(0, index + 1) } }"
                        :icon="index === stack.length - 1 ? 'text_snippet' : 'topic'"
                    />
                    <q-breadcrumbs-el :label="item" v-else/>
                  </template>
                </q-breadcrumbs>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <q-dialog v-model="addBucketDialog" persistent>
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">Bucket Name</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input dense v-model="bucketName" autofocus @keyup.enter="handleAddBucket" />
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Cancel" v-close-popup />
          <q-btn flat label="Add Bucket" @click="handleAddBucket" />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <q-dialog v-model="addPairDialog" persistent>
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">Add Pair</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input label="Key" dense v-model="pair.key" autofocus />
          <q-input class="q-mt-md" label="Value" dense v-model="pair.value" type="textarea" />
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Cancel" v-close-popup />
          <q-btn flat label="Add Pair" @click="handleAddPair" />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <div class="row items-start">
      <div class="col-10 offset-1">
        <div class="q-pa-sm">
          <q-table
              ref="entriesTable"
              virtual-scroll
              :rows="items"
              @request="onRequest"
              :filter="filter"
              row-key="id"
              hide-header
              :rows-per-page-options="[0]"
              :loading="loading"
              :columns="[{ name: 'name', label: 'Name', field: 'name', align: 'left', sortable: true }]"
              class="items-table"
              :virtual-scroll-sticky-size-start="48"
              hide-pagination
          >
            <template v-slot:top>
              <div class="full-width" style="display: inline-flex; align-items: center;">
                <h5 class="q-my-xs text-subtitle1">Contents</h5>
                <div style="margin-left: 15px; margin-right: 5px">
                  <q-fab
                      label=""
                      color="primary"
                      icon="add"
                      direction="right"
                      padding="5px"
                  >
                    <q-fab-action @click="addBucketDialog=true" color="purple" padding="5px" label="Bucket" />
                    <q-fab-action v-if="stack.length > 0" @click="addPairDialog=true" color="secondary" padding="5px" label="Pair" />
                  </q-fab>
                </div>

                <q-space/>

                <q-input class="q-my-xs" outlined placeholder="Search here..." dense autofocus debounce="300" color="primary" v-model="filter">
                  <template v-slot:append>
                    <q-icon name="search"/>
                  </template>
                </q-input>
              </div>
            </template>

            <template v-slot:body-cell-name="props">
              <q-td key="name" class="table-row" :props="props" @click="handleRowClick(props.row)">
                <entry-row :stack="stack" :entry="props.row" @refresh="refresh" />
              </q-td>
            </template>
          </q-table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {onMounted, ref, watch} from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {entries} from "@/store";
import EntryRow from "@/components/EntryRow";
import {useQuasar} from "quasar";

export default {
  components: {EntryRow},
  setup () {
    const entriesTable = ref(null)

    const stack = ref([])
    const filter = ref('')
    const loading = ref(true)

    const addBucketDialog = ref(false)
    const addPairDialog = ref(false)
    const bucketName = ref('')
    const pair = ref({key: "", value: ""})

    const store = entries()
    const items = ref([])

    const route = useRoute()
    const router = useRouter()

    const $q = useQuasar()

    // fetch the stack from the route params
    watch(
        () => route.params.stack,
        async stackFromParams => {
          console.log(stackFromParams)
          if (!stackFromParams) {
            stackFromParams = []
          }
          stack.value = stackFromParams.map(item => decodeURI(item))
          filter.value = ''

          if (entriesTable.value) {
            entriesTable.value.requestServerInteraction()
          }
        },
        { immediate: true }
    )

    onMounted(() => {
      refresh()
    })

    function refresh() {
      entriesTable.value.requestServerInteraction()
    }

    function onRequest(props) {
      const filter = props.filter
      // fetch data from "server"
      fetchFromServer(filter)
    }

    async function fetchFromServer (filter) {
      const response = await store.getEntries({
        level_stack: stack.value || [],
        filter: filter,
      })

      if (!response) {
        items.value = []
        return
      }

      items.value = response.map((item, index) => {
        return {
          id: index,
          name: item.name,
          type: item.type,
          content: item.value,
          original_content: item.value,
          original_name: item.name,
          is_bucket: item.is_bucket,
          child_buckets_count: item.no_of_child_bkts-1 || 0,
          child_pairs_count: item.no_of_pairs || 0,
          expanded: false,
          edit_name: false,
          edit_content: false,
        }
      })

      loading.value = false
      entriesTable.value.scrollTo(0, '-force')
    }

    function handleAddBucket() {
      if (!bucketName.value) {
        $q.notify({
          message: 'Bucket name cannot be empty',
          color: 'negative',
          icon: 'warning',
        })
        return
      }

      store.addBuckets({
        level_stack: stack.value,
        Buckets: [bucketName.value],
      }).then(() => {
        $q.notify({
          message: 'Bucket added successfully',
          color: 'positive',
          icon: 'done',
        })
        bucketName.value = ''
        addBucketDialog.value = false
        refresh()
      })
    }

    function handleAddPair() {
      if (!pair.value.key || !pair.value.value) {
        $q.notify({
          message: 'Key and Value cannot be empty',
          color: 'negative',
          icon: 'warning',
        })
        return
      }

      if (!isJsonString(pair.value.value)) {
        $q.notify({
          message: 'Value must be a valid JSON string',
          color: 'negative',
          icon: 'warning',
        })
        return
      }

      store.addPairs({
        level_stack: stack.value,
        Pairs: [{key: pair.value.key, value: JSON.parse(pair.value.value)}],
      }).then(() => {
        $q.notify({
          message: 'Pair added successfully',
          color: 'positive',
          icon: 'done',
        })
        pair.value = {key: "", value: ""}
        addPairDialog.value = false
        refresh()
      })
    }

    function isJsonString(str) {
      try {
        JSON.parse(str);
      } catch (e) {
        return false;
      }
      return true;
    }

    function handleRowClick(row) {
      if (!row.is_bucket) {
        row.expanded = !row.expanded
        return
      }

      if (!stack.value.includes(row.name)) {
        stack.value.push(row.name)
      }

      row.loading = true

      router.push({
        name: 'home',
        params: {
          stack: stack.value.map(item => encodeURI(item))
        }
      })
    }

    return {
      items,
      stack,

      //table
      entriesTable,
      filter,
      loading,
      onRequest,
      handleRowClick,

      // new bucket/pair
      addBucketDialog,
      addPairDialog,
      bucketName,
      pair,

      // dialog states

      handleAddBucket,
      handleAddPair,
      refresh,
    }
  },
}
</script>
