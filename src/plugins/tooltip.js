import { addClass, removeClass } from "@/util/helpers";

export default {
  install(app) {
    app.directive("tooltip", {
      mounted(el, binding) {
        // console.log(el, binding);
        const span = document.createElement("SPAN");
        const text = document.createTextNode(
          `Seats available: ${binding.value}`
        );
        span.appendChild(text);
        addClass(span, "tooltip");
        el.appendChild(span);

        const div = el.childNodes[0];
        function tooltipShow() {
          addClass(span, "tooltip-show");
        }
        div.addEventListener("mouseover", tooltipShow);

        function tooltipHide() {
          removeClass(span, "tooltip-show");
        }
        div.addEventListener("mouseout", tooltipHide);
      }
    });
  }
};
