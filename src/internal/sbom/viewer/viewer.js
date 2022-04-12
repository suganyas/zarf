const sbomSelector = document.getElementById('sbom-selector')
const distroInfo = document.getElementById('distro-info')
const modal = document.getElementById('modal')
const modalFader = document.getElementById('modal-fader')
const modalTitle = document.getElementById('modal-title')
const modalContent = document.getElementById('modal-content')
const artifactsTable = document.createElement('table')

document.body.appendChild(artifactsTable)

function initSelector(payload) {
    Object.keys(payload).sort().forEach(image => {
        sbomSelector.add(new Option(image));
    });

    sbomSelector.onchange();
}

function initData(selectedImage) {
    const payload = importedJSONData[selectedImage]

    const transformedData = payload.artifacts.map(artifact => {
        return [
            artifact.type,
            artifact.name,
            artifact.version,
            fileList(artifact.metadata),
            artifact.metadata.description || '-',
            (artifact.metadata.maintainer || '-').replace(/\u003c(.*)\u003e/, '&nbsp;|&nbsp;&nbsp;<a href="mailto:$1">$1</a>'),
            artifact.metadata.installedSize || '-',
        ];
    });

    const data = {
        "headings": [
            "Type",
            "Name",
            "Version",
            "Files",
            "Notes",
            "Maintainer",
            "Size",
        ],
        "data": transformedData,
    }

    if (window.dt) {
        window.dt.destroy()
    }

    distroInfo.innerHTML = payload.distro.prettyName

    window.dt = new simpleDatatables.DataTable(artifactsTable, {
        data,
        perPage: 20,
    })

}

function fileList(metadata) {
    const list = (metadata.files || [])
        .map(file => {
            return file.path || ''
        })
        .filter(test => test)

    if (list.length > 0) {
        flatList = list.sort().join('<br>');
        console.log(flatList);
        return `<a href="#" onClick="showModal('${metadata.package}','${flatList}')">${list.length} files</a>`
    }

    return '-';
}

function showModal(title, list) {
    modalTitle.innerText = `Files for ${title}`
    modalContent.innerHTML = list
    modalFader.className = "active";
    modal.className = "active";
}

function hideModal() {
    modalFader.className = ""
    modal.className = ""
    modalTitle.innerText = ""
    modalContent.innerHTML = ""
}

initSelector(importedJSONData)