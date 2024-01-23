import {setupTheme} from '../utils/theme.js'
import {setupUserNav} from '../components/user-nav.js'
import Loader from '../components/loader.js'

setupTheme()
setupUserNav()

const deleteBtn = document.getElementById('delete-btn')

deleteBtn && deleteBtn.addEventListener('click', async () => {
    if (confirm('Are you sure?, this action cannot be undone')) {
            Loader(true)

            const res = await fetch(location.href, {
                method: 'DELETE'
            })

            Loader()

            if (res.ok) {
                if (confirm('Delete successful, do you want to Logout as well?')) location.href = '/auth/logout'
                else location.reload()
            } else {
                const { apiError } = await res.json()
                confirm(apiError.description+ ', Do you want to refresh the page?') && location.reload()
        }
    }
})
