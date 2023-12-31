import {setupTheme} from '../utils/theme.js'
import {setupUserNav} from '../components/user-nav.js'
import {doFile} from '../components/upload.js'

setupTheme()
setupUserNav()

const uploadArea = document.querySelector('.upload-area')
uploadArea.setAttribute('open', '')

doFile(async (file, metadata) => {
    const res = await fetch(location.href, {
        method: 'PUT',
        body: file,
        headers: {
            'Content-Type': file.type,
            'File-Private-Url-Expires': metadata.xPrivateUrlExpires,
            'File-Is-Public': metadata.isPublic,
            'File-Auto-Deleted-At': metadata.autoDeletedAt
        }
    })

    if (res.ok) {
        alert('Updated File successfully')
        location.reload()
    } else {
        const {errorDescription} = await res.json()
        confirm(errorDescription+ ' Do you want to refresh the page?') && location.reload()
    }
})

{
    const deleteBtn = document.querySelector('.delete-btn')
    deleteBtn.addEventListener('click', async () => {
        if (confirm('Are you sure?, this action cannot be undone')) {
            const res = await fetch(location.href, {
                method: 'DELETE'
            })

            if (res.ok) {
                const username = location.pathname.split('/')[2]
                location.href = '/dashboard/' + username
            } else {
                const {errorDescription} = await res.json()
                confirm(errorDescription+ ', Do you want to refresh the page?') && location.reload()
            }
        }
    })
}
