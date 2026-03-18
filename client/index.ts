import { registerHoverCapabilityController } from "@client/hover";
import { ArchiveCarousel } from "./archive-carousel";
import { me } from "@client/inline";

if (window !== undefined) {
    // inline.ts exports
    (window as any).me = me;

    // archive-carousel.ts exports
    (window as any).ArchiveCarousel = ArchiveCarousel;
}

registerHoverCapabilityController();
