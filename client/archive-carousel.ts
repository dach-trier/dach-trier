// prettier-ignore

export function ArchiveCarousel(element: HTMLElement) {
    let   active:     HTMLElement = element.querySelector(".archive__carousel__slide--active")!;
    const prevButton: HTMLElement = element.querySelector(".archive__carousel__control__prev")!;
    const nextButton: HTMLElement = element.querySelector(".archive__carousel__control__next")!;
    const timer:      HTMLElement = element.querySelector(".archive__carousel__timer")!;
    let   start:      number      = performance.now();
    const duration:   number      = 5000;

    function getFirstChild(element: HTMLElement): HTMLElement {
        return element.children[0]! as HTMLElement;
    }

    function getLastChild(element: HTMLElement): HTMLElement {
        const  children = element.children;
        const nchildren = element.children.length;

        return children[nchildren - 1]! as HTMLElement;
    }

    function getPreviousSibling(element: HTMLElement): HTMLElement {
        if (element.previousElementSibling)
            return element.previousElementSibling as HTMLElement;

        return getLastChild(element.parentElement! as HTMLElement);
    }

    function getNextSibling(element: HTMLElement): HTMLElement {
        if (element.nextElementSibling)
            return element.nextElementSibling as HTMLElement;

        return getFirstChild(element.parentElement! as HTMLElement);
    }

    function selectPrevSibling() {
        start = performance.now();
        active.classList.remove("archive__carousel__slide--active");
        const previous = getPreviousSibling(active);
        active = previous;
        active.classList.add("archive__carousel__slide--active");
    }

    function selectNextSibling() {
        start = performance.now();
        active.classList.remove("archive__carousel__slide--active");
        const next = getNextSibling(active);
        active = next;
        active.classList.add("archive__carousel__slide--active");
    }

    function tick(time: number) {
        const elapsed = time - start;

        if (elapsed >= duration) {
            selectNextSibling();
            tick(start);
            return;
        }

        timer.style.width = `${elapsed / duration * 100}%`;
        requestAnimationFrame(tick);
    }

    requestAnimationFrame(tick);

    prevButton.addEventListener("click", selectPrevSibling);
    nextButton.addEventListener("click", selectNextSibling);
}
