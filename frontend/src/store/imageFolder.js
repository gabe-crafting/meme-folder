import { writable } from 'svelte/store';

export const imageFolderPath = writable(localStorage.getItem("folderPath") || "");

export function updateFolderPath(newVal) {
    imageFolderPath.set(newVal)
    localStorage.setItem("folderPath", newVal);
}