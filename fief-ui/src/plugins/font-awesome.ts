import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faLock,
  faUser,
  faUserPlus,
} from "@fortawesome/free-solid-svg-icons";

library.add(faUser, faUserPlus, faLock);

export { FontAwesomeIcon };