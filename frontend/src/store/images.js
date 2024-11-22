import { writable } from "svelte/store";

export const images = writable([])
export const selectedTags = writable([])
export const allTags = writable([])

export function updateImages(res) {
    images.set(res)
    // Object.keys(imagesTemp.reduce((acc, image)=>{image.tags.forEach(tag=>{acc[tag] = true}); return acc}, {}))
    // TODO: or redo with flat map
    allTags.set([...new Set(res.map((image) => image.Tags).flat())]);
}