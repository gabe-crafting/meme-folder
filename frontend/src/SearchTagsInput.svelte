<script lang="ts">
    import { onMount } from "svelte";
    import { selectedTags } from "./store/images";

    export let allTags = [];
    let tags = [];
    let input = "";
    let matchedTags = [];
    let selectedIndex = 0;

    selectedTags.subscribe((value) => {
        tags = value;
    });

    function addTagToSearch() {
        const selectedTag = matchedTags[selectedIndex];
        if (selectedTag != null && !tags.includes(selectedTag)) {
            selectedTags.set([...tags, matchedTags[selectedIndex]]);
            input = "";
            selectedIndex = 0;
        }
    }

    function removeTagFromSearch(tag) {
        selectedTags.set(tags.filter((el) => el !== tag));
    }

    function onkeydownHandler(event) {
        if (event.key === "Enter") {
            addTagToSearch();
        }

        if (event.key === "Backspace" && input === "") {
            selectedTags.set(tags.slice(0, -1));
        }

        if (event.key === "ArrowDown") {
            event.preventDefault();
            if (selectedIndex + 1 < matchedTags.length) {
                selectedIndex += 1;
            }
        }

        if (event.key === "ArrowUp") {
            event.preventDefault();
            if (selectedIndex - 1 >= 0) {
                selectedIndex -= 1;
            }
        }
    }

    function filterSpaces() {
        input = input.replace(/\s+/g, "").toLowerCase();
        matchedTags = allTags.filter(
            (tag) => tag.includes(input) && !tags.includes(tag),
        );
    }

    onMount(() => {
        const searchTagInput = document.getElementById("searchTagInput");

        searchTagInput.focus();
    });
</script>

<div class="search-tags-container">
    <div class="search-tags-input-container">
        {#each tags as tag}
            <span class="tag" on:mousedown={() => removeTagFromSearch(tag)}>
                {tag}
            </span>
        {/each}
        <input
            id="searchTagInput"
            bind:value={input}
            on:input={filterSpaces}
            on:keydown={onkeydownHandler}
        />
    </div>
    {#if matchedTags.length > 0 && input != ""}
        <div class="matched-tags-container">
            <ul>
                {#each matchedTags as tag}
                    <li class:selected={matchedTags[selectedIndex] === tag}>
                        {tag}
                    </li>
                {/each}
            </ul>
        </div>
    {/if}
</div>

<style>
    .search-tags-container {
        position: relative;
    }
    .matched-tags-container {
        z-index: 3;
        width: 100%;
        position: absolute;
        background-color: black;
        color: white;
    }
    .matched-tags-container ul {
        list-style-type: none;
        padding: 0;
        margin: 0;
    }

    .search-tags-input-container {
        background-color: white;
        border: solid 1px gray;
        margin: 5px;
        display: flex;
        flex-wrap: wrap;
    }

    .search-tags-input-container .tag {
        background-color: gray;
        color: white;
        margin: 2px;
        padding: 2px 5px;
        border-radius: 5px;
        cursor: pointer;
    }

    .search-tags-input-container input {
        border: none;
        flex-grow: 1;
        width: auto;
    }

    .matched-tags-container li.selected {
        background-color: white;
        color: black;
    }
</style>
