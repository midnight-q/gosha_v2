import 'package:get/get_state_manager/src/simple/get_controllers.dart';
import 'package:get/get.dart';
import 'package:gosha_app/sdk/model.dart';

class ModelController extends GetxController {
  var isShowFilters = false.obs;
  var isShowServiceModel = false.obs;
  var searchQuery = "".obs;
  var models = List<Model>.empty().obs;
  void loadModels(String path) async{
    var res = await modelFind(path);
    this.models.value = res.theList;
  }
}
