<script>
    import { onMount } from "svelte";
    import { GetImages, AddImageByLink } from "../wailsjs/go/main/App.js";
    import Image from "./Image.svelte";
    import {
        allTags,
        images,
        selectedTags,
        updateImages,
    } from "./store/images.js";
    import SearchTagsInput from "./SearchTagsInput.svelte";

    export let imageFolder;

    let photos = [];
    let selectedTagsLocal = [];
    let tempAllTags = [];
    let tempPhotos = [];

    export async function addNewImage(link, newTag) {
        const imagesTemp = Object.values(
            await AddImageByLink(imageFolder, link, newTag),
        );
        updateImages(imagesTemp);
    }

    async function fetchImages() {
        const imagesTemp = Object.values(await GetImages(imageFolder));
        updateImages(imagesTemp);
    }

    images.subscribe((value) => {
        photos = value;
    });

    selectedTags.subscribe((value) => {
        selectedTagsLocal = value;
    });

    allTags.subscribe((value) => {
        tempAllTags = value;
    });

    onMount(async () => {
        await fetchImages();
    });

    $: if (imageFolder) {
        fetchImages();
    }

    $: if (selectedTagsLocal) {
        tempPhotos = photos.filter((photo) => {
            return selectedTagsLocal.every((selectedTag) =>
                photo.Tags.includes(selectedTag),
            );
        });

        if (tempPhotos.length === 0) {
            tempPhotos = photos;
        }
    }
</script>

<div class="gallery-container">
    <SearchTagsInput allTags={tempAllTags} />
    <div class="gallery">
        <!-- svelte-ignore a11y-img-redundant-alt -->
        {#each tempPhotos as photo}
            <Image image={photo} />
        {/each}
    </div>
</div>

<style>
    .gallery {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap;
    }
</style>
