import {setupUserNav} from '../components/user-nav.js'
import {setupTheme} from '../utils/theme.js'
import loader from '../components/loader.js'
import {doFile} from '../components/upload.js'
import Loader from '../components/loader.js'

setupUserNav()
setupTheme()

doFile(async (file, metadata) => {
    Loader(true)

    const res = await fetch(location.href, {
        method: 'POST',
        headers: {
            'Content-Type': file.type,
            'File-Auto-Delete-At': metadata.autoDeleteAt,
            'File-Is-Public': metadata.isPublic,
            'File-Name': file.name,
            'File-Private-Url-Expires': metadata.privateUrlExpires,
        },
        body: file,
    })

    Loader()
    if (res.ok) {
        alert('Upload successfully')
        location.reload()
    } else {
        const {errorDescription} = await res.json()
        confirm(errorDescription+ ', Do you want to refresh the page?') && location.reload()
    }

})
