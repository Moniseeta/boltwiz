import EntriesList from "@/pages/EntriesList";

const routes = [
    {
        name: 'home',
        path: '/:stack*',
        component: EntriesList
    },
]

export default routes