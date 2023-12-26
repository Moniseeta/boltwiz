<template>
  <q-layout view="lHh Lpr lFf">
    <div class="main-content">
      <q-header>
        <q-toolbar class="q-py-sm">
          <div class="offset-1 col-9">
            <div class="row" style="align-items: center; font-size: large">
              <q-icon size="xl" class="q-mr-sm" name="img:img/icon.png" />
                BoltWiZ
            </div>
          </div>
        </q-toolbar>
      </q-header>

      <q-page-container>
        <router-view></router-view>
      </q-page-container>

      <div class="footer">
        <div class="q-py-md text-center">
          <p class="q-my-none">Reach us @ <a style="color: white" href="https://github.com/Moniseeta/boltdbgui"> Github </a></p>
        </div>
      </div>
    </div>
  </q-layout>
</template>

<script>

export default {
  name: 'AppLayout',
  data() {
    return {}
  },

}
</script>
