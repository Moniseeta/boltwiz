<template>
  <div class="row">
    <q-icon name="topic" v-if="row.is_bucket" class="q-mt-sm q-mr-sm" color="indigo-3" size="sm"/>
    <q-icon name="text_snippet" v-else class="q-mt-sm q-mr-sm" color="grey-3" size="sm"/>
    <template v-if="!row.editing">
      <div class="row" style="width: calc(100% - 35px)">
        <div class="editable-text">
          <div style="line-height: 40px !important;" >{{ row.name }}
            <q-badge class="q-ml-sm" color="purple-4" v-if="row.child_buckets_count > 0">
              <q-icon name="topic" color="white" class="q-mr-xs" /> {{ row.child_buckets_count }}
            </q-badge>
            <q-badge class="q-ml-sm" color="light-green-4" v-if="row.child_pairs_count > 0">
              <q-icon name="text_snippet" color="white" class="q-mr-xs" /> {{ row.child_pairs_count }}
            </q-badge>
            <q-spinner-ios class="q-ml-sm" v-if="row.loading" size="sm" color="grey" />
          </div>
          <div class="actions">
            <q-btn v-if="!row.is_bucket" @click.stop="editItem($event, row)" class="edit-btn" color="grey-6" round dense flat icon="edit"/>
            <q-btn @click.stop="deleteEntry($event, row)" class="edit-btn" color="red-4" round dense flat icon="delete"/>
          </div>
        </div>
      </div>
    </template>

    <template v-else>
      <q-input @click.stop @keydown.enter="renameEntry($event, row)" dense debounce="300" color="primary" style="width: calc(100% - 35px)" v-model="row.name">
        <template v-slot:append>
          <q-btn @click.stop="renameEntry($event, row)" round dense flat color="blue" icon="check_circle"/>
          <q-btn @click.stop="row.name = row.original_name; row.editing = false" round dense flat icon="cancel"/>
        </template>
      </q-input>
    </template>

    <template v-if="!row.is_bucket">
      <div @click.stop style="cursor:text;" :class="!row.expanded? 'row-content-container' : 'row-content-container expanded'">
        <div v-if="!row.edit_content" class="row-content" style="padding: 10px">
          <div class="row-content-text">{{ formatContent(row) }}</div>
          <q-btn @click.stop="row.format = !row.format" round dense flat class="no-hover q-mr-sm" color="blue" icon="code"/>
          <q-btn @click.stop="copy(row)" round dense flat class="no-hover q-mr-sm" color="blue" icon="content_copy"/>
          <q-btn @click.stop="row.edit_content = true" round dense flat class="no-hover" color="blue" icon="edit"/>
        </div>
        <div v-else class="row-content">
          <q-input type="textarea" dense debounce="300" color="primary" style="width: 100%" v-model="row.content">
            <template v-slot:append>
              <q-btn round dense flat color="blue" @click.stop="updateEntryValue($event, row)" icon="check_circle"/>
              <q-btn @click.stop="row.edit_content = false; row.content = row.original_content" round dense flat icon="cancel"/>
            </template>
          </q-input>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import {toRefs} from 'vue'
import { useQuasar } from 'quasar'
import useClipboard from 'vue-clipboard3'
import {entries} from "@/store";

export default {
  props: {
    entry: {
      type: Object,
      required: true
    },
    stack: {
      type: Array,
      required: true
    }
  },
  emits: [
      'refresh',
  ],
  setup (props, { emit }) {
    const { entry, stack } = toRefs(props)

    const store = entries()

    const $q = useQuasar()

    const { toClipboard } = useClipboard()

    function formatContent(row) {
      if (row.format) {
        return JSON.stringify(JSON.parse(row.content), null, 4)
      }
      return row.content
    }

    const copy = async (row) => {
      const content = formatContent(row)
      try {
        await toClipboard(content)
        $q.notify({
          message: 'Copied to clipboard',
          color: 'grey',
          textColor: 'black',
          icon: 'check',
          position: 'bottom',
          timeout: 100,
        })
      } catch (e) {
        console.error(e)
      }
    }

    function isJsonString(str) {
      try {
        JSON.parse(str);
      } catch (e) {
        return false;
      }
      return true;
    }

    async function deleteEntry(_, row) {
      $q.dialog({
        title: 'Confirm',
        message: 'Are you sure you want to delete this entry?',
        cancel: true,
        persistent: true
      }).onOk(() => {
        const request = {
          level_stack: stack.value,
          key: row.name
        }
        store.deleteEntry(request).then(() => {
          $q.notify({
            message: 'Entry deleted',
            color: 'grey',
            textColor: 'black',
            icon: 'check',
            position: 'bottom',
            timeout: 100,
          })
        }).catch(() => {
          $q.notify({
            message: 'Error deleting entry',
            color: 'red',
            textColor: 'white',
            icon: 'error',
            position: 'bottom',
            timeout: 100,
          })
        }).finally(() => {
          emit('refresh')
        })
      })
    }

    function renameEntry(_, row) {
      if (!row.name) {
        $q.notify({
          message: 'Name cannot be empty',
          color: 'negative',
          icon: 'warning',
        })
        return
      }

      $q.dialog({
        title: 'Confirm',
        message: 'Are you sure you want to rename this entry?',
        cancel: true,
        persistent: true
      }).onOk(() => {
        const request = {
          level_stack: stack.value,
          key: row.original_name,
          new_key: row.name,
        }
        store.renameKey(request).then(() => {
          $q.notify({
            message: 'Entry deleted',
            color: 'grey',
            textColor: 'black',
            icon: 'check',
            position: 'bottom',
            timeout: 100,
          })
        }).catch(() => {
          $q.notify({
            message: 'Error deleting entry',
            color: 'red',
            textColor: 'white',
            icon: 'error',
            position: 'bottom',
            timeout: 100,
          })
        }).finally(() => {
          emit('refresh')
        })
      })
    }

    function updateEntryValue(_, row) {
      if (!isJsonString(row.content)) {
        $q.notify({
          message: 'Value must be a valid JSON string',
          color: 'negative',
          icon: 'warning',
        })
        return
      }

      const request = {
        level_stack: stack.value,
        key: row.name,
        new_value: JSON.parse(row.content)
      }

      store.updateValue(request).then(() => {
        $q.notify({
          message: 'Entry updated',
          color: 'grey',
          textColor: 'black',
          icon: 'check',
          position: 'bottom',
          timeout: 100,
        })
      }).catch(() => {
        $q.notify({
          message: 'Error updating entry',
          color: 'red',
          textColor: 'white',
          icon: 'error',
          position: 'bottom',
          timeout: 100,
        })
      }).finally(() => {
        row.edit_content = false
      })
    }

    return {
      row: entry,

      formatContent,
      editItem(event, item) {
        item.editing = !item.editing
      },

      copy,
      updateEntryValue,
      deleteEntry,
      renameEntry
    }
  },
}
</script>
