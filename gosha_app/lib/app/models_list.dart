import 'package:flutter/material.dart';
import 'package:flutter_staggered_grid_view/flutter_staggered_grid_view.dart';
import 'package:get/get.dart';
import 'package:gosha_app/controller/app_controller.dart';
import 'package:gosha_app/controller/model_controller.dart';

class ModelList extends StatelessWidget {
  ModelList({Key? key}) : super(key: key);

  ModelController modelC = Get.put(ModelController());
  AppController applicationC = Get.find();

  @override
  Widget build(BuildContext context) {
    return Obx(() {
      modelC.loadModels(applicationC.currentPath.value);
      return Padding(
        padding: EdgeInsets.all(4),
        child: StaggeredGridView.countBuilder(
          crossAxisCount: 4,
          itemBuilder: (BuildContext context, int index) {
            return Text("${modelC.models[index].Name}");
          },
          staggeredTileBuilder: (int index) {
            return StaggeredTile.count(1, 1);
          },
          itemCount: modelC.models.length,
        ),
      );
    });
  }
}
