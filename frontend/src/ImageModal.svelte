<script lang="ts">
    import { AddTag, RemoveTag } from "../wailsjs/go/main/App.js";
    import { images, updateImages } from "./store/images";
    import { imageFolderPath } from "./store/imageFolder";
    import { get } from "svelte/store";
    import { onMount } from "svelte";

    export var close;
    export var image;
    let newTagName = "";

    function focusInput() {
        const newTagInput = document.getElementById("newTagInput");
        newTagInput.focus();
    }

    async function addTag() {
        updateImages(
            Object.values(
                await AddTag(get(imageFolderPath), image.Id, newTagName),
            ),
        );
        newTagName = "";
    }

    async function removeTag(removedTag) {
        updateImages(
            Object.values(
                await RemoveTag(get(imageFolderPath), image.Id, removedTag),
            ),
        );
        focusInput();
    }

    function newTagInputKeyDown(event) {
        if (event.key === "Enter") addTag();

        if (event.key === "Escape") close();
    }

    function filterSpaces() {
        newTagName = newTagName.replace(/\s+/g, "").toLowerCase();
    }

    onMount(() => {
        focusInput();
    });
</script>

<div class="modal-container">
    <div class="modal">
        <button class="modal-button" on:click={close}>Close modal</button>
        <div class="modal-image-container">
            <img src={image.Path} alt={image.Name} />
        </div>
        <div class="add-tag-container">
            <input
                bind:value={newTagName}
                on:input={filterSpaces}
                on:keydown={newTagInputKeyDown}
                placeholder="tag name"
                id="newTagInput"
            />
            <button class="button" on:click={addTag}>Add tag</button>
        </div>
        <div class="add-tag-container tag-container">
            {#each image.Tags as tag}
                <span class="tag" on:mousedown={() => removeTag(tag)}
                    >{tag}
                </span>
            {/each}
        </div>
    </div>
</div>

<style>
    .modal-image-container {
        max-height: 100%;
    }

    .modal-image-container img {
        max-height: 50vh;
    }

    .modal-container {
        z-index: 2;
        display: flex;
        justify-content: center;
        align-items: center;
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100vh;
        background: rgba(0, 0, 0, 0.8);
    }

    .modal-button {
        position: absolute;
        right: 5px;
        top: 5px;
    }

    .add-tag-container {
        width: 100%;
        margin-bottom: 20px;
    }

    .tag {
        background-color: #c9d7dd;
        padding: 2px 8px;
        margin: 2px;
        color: #637a9f;
        border-radius: 50px;
        cursor: pointer;
        border: none;
    }

    .modal {
        overflow: auto;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        position: relative;
        height: 90vh;
        background-color: rgba(27, 38, 54, 1);
        width: 90%;
    }
</style>
