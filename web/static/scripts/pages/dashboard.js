import {setupUserNav} from '../components/user-nav.js'
import {setupTheme} from '../utils/theme.js'
import {doFile} from '../components/upload.js'

setupUserNav()
setupTheme()

doFile(async (file, metadata) => {
    const res = await fetch(location.href, {
        method: 'POST',
        headers: {
            'Content-Type': file.type,
            'File-Auto-Deleted-At': metadata.autoDeletedAt,
            'File-Is-Public': metadata.isPublic,
            'File-Name': file.name,
            'File-Private-Url-Expires': metadata.xPrivateUrlExpires,
        },
        body: file,
    })


    if (res.ok) {
        alert('Upload successfully')
        location.reload()
    } else {
        const {errorDescription} = await res.json()
        confirm(errorDescription+ ', Do you want to refresh the page?') && location.reload()
    }
})